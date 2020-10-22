package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/isutare412/MukGo/server"
	pb "github.com/isutare412/MukGo/server/api/proto"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
	"github.com/streadway/amqp"
)

// Server runs as API server for MukGo service. Server should be created with
// NewServer function.
type Server struct {
	mux  *http.ServeMux
	mqss *mq.Session
	hc   *http.Client

	packetChans *server.ChannelMap
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

	// http client for authorization
	s.hc = &http.Client{}

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

	// create response packet map
	s.packetChans = server.NewChannelMap()

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

	// retrieve packet channel
	ch := s.packetChans.Pop(d.CorrelationId)
	if ch == nil {
		return true, nil // handler dropped by timeout
	}
	defer close(ch)

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

	case server.PTDAUser:
		var p server.DAPacketUser
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDARestaurants:
		var p server.DAPacketRestaurants
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDAReviews:
		var p server.DAPacketReviews
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	default:
		parseErr = fmt.Errorf("no parser for %d", int(packetType))
	}

	// on packet parsing failed
	if parseErr != nil {
		return false, fmt.Errorf("on DBResponse: %v", parseErr)
	}

	// send parsed packet
	ch <- packet

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
	s.mux.HandleFunc("/reviews", s.handleReviews)
	s.mux.HandleFunc("/restaurant", s.handleRestaurant)
	s.mux.HandleFunc("/restaurants", s.handleRestaurants)
}

// send2DB send packet to database server. It returns chan Packet as response.
// Response packet from database server is passed through the channel when
// api server receives the response packet. In error case or timeout,
// returned channel is closed. So it is safe to wait for the channel.
func (s *Server) send2DB(
	p server.Packet,
) (<-chan server.Packet, error) {
	// wrapper with done channel when response is called
	pch := make(chan server.Packet)

	// register response handler
	correlationID := <-s.packetChans.IDGet
	if err := s.packetChans.Register(correlationID, pch); err != nil {
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
		close(s.packetChans.Pop(correlationID))
		return nil, fmt.Errorf("on send2DB: %v", err)
	}

	// set response handler timeout
	go func(corrID string) {
		<-time.After(responseTimeout)

		ch := s.packetChans.Pop(corrID)
		if ch == nil {
			return
		}
		close(ch)
	}(correlationID)

	return pch, nil
}

func (s *Server) authenticate(h http.Header) (uid, name string, err error) {
	const (
		authKey = "Authorization"
		apiURL  = "https://www.googleapis.com/oauth2/v3/userinfo"
	)

	auth := h.Get(authKey)
	if auth == "" {
		err = fmt.Errorf("No Authorization header")
		return
	}

	// extract token from header
	tokens := strings.SplitN(auth, " ", 2)
	if len(tokens) < 2 {
		err = fmt.Errorf("Authorization header should be space seperated")
		return
	}

	tokenType, accessToken := tokens[0], tokens[1]

	// easy authorization for develeopment purpose
	if tokenType == "Mukgoer" {
		uid = accessToken
		name = accessToken
		return
	}

	// build request struct
	req, terr := http.NewRequest("GET", apiURL, nil)
	if terr != nil {
		err = fmt.Errorf("failed request: %v", terr)
		return
	}

	// send to google auth api
	req.Header.Set(authKey, "Bearer "+accessToken)
	res, terr := s.hc.Do(req)
	if terr != nil {
		err = fmt.Errorf("failed send: %v", terr)
		return
	}

	// decode json
	var uc UserClaim
	terr = json.NewDecoder(res.Body).Decode(&uc)
	if terr != nil {
		err = fmt.Errorf("failed decode: %v", terr)
		return
	}

	if uc.Sub == "" {
		err = fmt.Errorf("invalid access token(%v)", accessToken)
		return
	}

	uid = uc.Sub
	name = uc.Name
	return
}

func baseHeader(h http.Header) {
	h.Set("Content-Type", "application/json; charset=utf-8")
}

// httpError responses to client with proper http error message.
func httpError(w http.ResponseWriter, errno int, code pb.Code) {

	// write custom error code to body
	baseHeader(w.Header())
	w.WriteHeader(errno)
	reason := pb.ErrorReason{Code: code}
	ser, err := json.Marshal(&reason)
	if err != nil {
		panic(fmt.Errorf("on httpError: %v", err))
	}
	w.Write(ser)
}

// getError checks if p is error packet, then returns its ErrorType.
func getError(p server.Packet) server.ErrorType {
	ep, ok := p.(*server.DAPacketError)
	if ok {
		return ep.ErrorType
	}
	return server.ETInvalid
}
