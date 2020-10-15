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

const responseTimeout = 6 * time.Second

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
	console.Info("connect to RabbitMQ...")
	if err := mqSession.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	s.mqss = mqSession
	console.Info("session(%q) established between RabbitMQ", mqaddr)

	// create ResponseMux
	s.handles = server.NewHandleMap()

	// start to listen from RabbitMQ
	err := s.mqss.Consume(server.MGDB, server.DB2API, s.onDBResponse, 2)
	if err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}

	// addlitionaly send logs to RabbitMQ.
	console.AddLogHandler(
		func(l console.Level, format string, v ...interface{}) bool {
			packet := server.PacketLog{
				Timestamp: time.Now(),
				LogLevel:  l,
				Msg:       fmt.Sprintf(format, v...),
			}

			if err := s.mqss.Publish(
				server.MGLogs,
				"",
				server.API,
				&packet,
			); err != nil {
				return false
			}
			return true
		},
	)

	return s, nil
}

func (s *Server) onDBResponse(d *amqp.Delivery) (bool, error) {
	_, packetType, err := mq.ParseHeader(d.Headers)
	if err != nil {
		return false, fmt.Errorf("on DBResponse: %v", err)
	}

	// retrieve handler
	handler := s.handles.Pop(d.CorrelationId)
	if handler == nil {
		return true, nil // handler dropped by timeout
	}

	var packet server.Packet
	var parseErr error

	// parse packet
	switch packetType {
	case server.PTDAAck:
		var p server.DAPacketAck
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDAError:
		var p server.DAPacketError
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDANoSuchUser:
		var p server.DAPacketNoSuchUser
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDANoSuchRestaurant:
		var p server.DAPacketNoSuchRestaurant
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDAUser:
		var p server.DAPacketUser
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDARestaurants:
		var p server.DAPacketRestaurants
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	default:
		parseErr = fmt.Errorf("no parser for %d", int(packetType))
	}

	// on packet parsing failed
	if parseErr != nil {
		go handler(false, nil)
		return false, fmt.Errorf("on DBResponse: %v", parseErr)
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
	s.mux.HandleFunc("/user", s.handleUser)
	s.mux.HandleFunc("/review", s.handleReview)
	s.mux.HandleFunc("/restaurant", s.handleRestaurant)
	s.mux.HandleFunc("/restaurants", s.handleRestaurants)
}

func (s *Server) send2DB(
	p server.Packet,
	response func(bool, server.Packet),
) (<-chan struct{}, error) {
	// wrapper with done channel when response is called
	done := make(chan struct{})
	wrapper := func(b bool, p server.Packet) {
		response(b, p)
		done <- struct{}{}
		close(done)
	}

	// register response handler
	correlationID := <-s.handles.IDGet
	if err := s.handles.Register(correlationID, wrapper); err != nil {
		return nil, fmt.Errorf("on send2DB: %v", err)
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
		return nil, fmt.Errorf("on send2DB: %v", err)
	}

	// set response handler timeout
	go func(corrID string) {
		<-time.After(responseTimeout)

		handler := s.handles.Pop(corrID)
		if handler == nil {
			return
		}
		handler(false, nil)
	}(correlationID)

	return done, nil
}

// httpError responses to client with proper http error message.
func httpError(w http.ResponseWriter, errno int) {
	http.Error(w, http.StatusText(errno), errno)
}
