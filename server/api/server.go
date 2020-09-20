package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
	"github.com/streadway/amqp"
)

// Server runs as API server for MukGo service. Server should be created with
// NewServer function.
type Server struct {
	mux  *http.ServeMux
	mqss *mq.Session

	handles *server.HandleMap
}

var baseConfig = &mq.SessionConfig{
	Exchanges: map[string]mq.ExchangeConfig{
		server.MGLogs: {
			Name: server.MGLogs,
			Type: "fanout",
		},
		server.MGDB: {
			Name: server.MGDB,
			Type: "direct",
			Queues: map[string]mq.QueueConfig{
				server.API2DB: {
					Name:       server.API2DB,
					RouteKey:   server.API2DB,
					Durable:    true,
					AutoDelete: false,
				},
				server.DB2API: {
					Name:       server.DB2API,
					RouteKey:   server.DB2API,
					Durable:    true,
					AutoDelete: false,
				},
			},
		},
	},
}

// NewServer creates Server struct safely.
func NewServer(cfg *ServerConfig) (*Server, error) {
	var s = &Server{
		mux: http.NewServeMux(),
	}

	// register api handlers
	s.registerHandlers()
	console.Info("registered handlers")

	// establish rabbitmq session
	mqaddr := fmt.Sprintf("%s:%d", cfg.RabbitMQ.IP, cfg.RabbitMQ.Port)
	baseConfig.User = cfg.RabbitMQ.User
	baseConfig.Password = cfg.RabbitMQ.Password
	baseConfig.Addr = mqaddr
	mqSession := mq.NewSession("api", baseConfig)

	// connecte the session
	if err := mqSession.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	s.mqss = mqSession
	console.Info("session(%q) established between RabbitMQ", mqaddr)

	// create ResponseMux
	s.handles = server.NewHandleMap()

	err := s.mqss.Consume(server.MGDB, server.DB2API, s.onDBResponse)
	if err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}

	return s, nil
}

func (s *Server) onDBResponse(d *amqp.Delivery) (bool, error) {
	header := d.Headers
	if header == nil {
		return false, fmt.Errorf("header does not exists in delievery")
	}
	msgType, ok := header[server.MsgType].(int)
	if !ok {
		return false, fmt.Errorf("invalid message type")
	}

	// retrieve handler
	handler := s.handles.Pop(d.CorrelationId)
	if handler == nil {
		return true, nil // handler dropped by timeout
	}

	var packet server.Packet
	var parseErr error

	// parse packet
	switch server.PacketType(msgType) {
	case server.PTAck:
		var p server.PacketAck
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)
	default:
		parseErr = fmt.Errorf("no parser for %d", msgType)
	}

	// on packet parsing failed
	if parseErr != nil {
		go handler(false, nil)
		return false, fmt.Errorf("onDBResponse: %v", parseErr)
	}

	// call handler
	go handler(true, packet)

	return true, nil
}

// ListenAndServe starts Server. If addr is blank, ":http" is used, which
// uses 80 port.
func (s *Server) ListenAndServe(addr string) error {
	if addr == "" {
		addr = ":http"
	}

	// Assign ServeMux to Server.
	httpserver := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	return httpserver.ListenAndServe()
}

func (s *Server) registerHandlers() {
	s.mux.HandleFunc("/devel", s.handlerDevel)
	s.mux.HandleFunc("/review", s.handlerReview)
}

func (s *Server) handlerReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		packet := server.PacketReview{}
		wait := make(chan struct{})

		s.send2DB(
			&packet,
			func(success bool, p server.Packet) {
				defer func() {
					wait <- struct{}{}
				}()

				if !success {
					console.Info("timeout!")
					http.Error(w, "", http.StatusInternalServerError)
					return
				}

				console.Info("received ACK!")
			},
		)

		<-wait
	}
}

func (s *Server) handlerDevel(w http.ResponseWriter, r *http.Request) {
	// Parse request into RestRequest.
	var req RestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		console.Warning("failed to decode request: %v", err)
		return
	}

	s.sendLog("message from %q: %q", req.User, req.Message)

	// marshal response into byte slice
	res := RestResponse{"Hello, Client!"}
	resBytes, err := json.Marshal(res)
	if err != nil {
		console.Warning("failed to encode response: %v", err)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}

func (s *Server) sendLog(format string, v ...interface{}) {
	packet := server.PacketLog{
		Timestamp: time.Now(),
		Msg:       fmt.Sprintf(format, v...),
	}

	if err := s.mqss.Publish(
		server.MGLogs,
		"",
		server.API,
		&packet,
	); err != nil {
		console.Error("failed to publish log: %v", err)
		return
	}
}

func (s *Server) send2DB(
	p server.Packet,
	response func(bool, server.Packet),
) error {
	// register response handler
	correlationID := <-s.handles.IDGet
	if err := s.handles.Register(correlationID, response); err != nil {
		return fmt.Errorf("on send2DB: %v", err)
	}

	// request RPC
	if err := s.mqss.RPC(
		server.MGDB,
		server.API2DB,
		server.API,
		server.DB2API,
		correlationID,
		p,
	); err != nil {
		s.handles.Pop(correlationID)
		return fmt.Errorf("on send2DB: %v", err)
	}

	// set response handler timeout
	go func(corrID string) {
		<-time.After(3 * time.Second)

		handler := s.handles.Pop(corrID)
		if handler == nil {
			return
		}
		handler(false, nil)
	}(correlationID)

	return nil
}
