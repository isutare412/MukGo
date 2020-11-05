package api

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/isutare412/MukGo/server"
	pb "github.com/isutare412/MukGo/server/api/proto"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
	"github.com/streadway/amqp"
	"google.golang.org/protobuf/proto"
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

const headerJSON = "application/json"
const headerProtobuf = "application/protobuf"

var puOption = proto.UnmarshalOptions{AllowPartial: true}

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

	case server.PTDAUsers:
		var p server.DAPacketUsers
		packet = &p
		parseErr = json.Unmarshal(d.Body, &p)

	case server.PTDARestaurant:
		var p server.DAPacketRestaurant
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
	s.mux.HandleFunc("/ranking", s.handleRanking)
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

// marshalQuery converts map[string][]string to map[string]string, keeping
// only the first one.
func marshalQuery(q url.Values) map[string]string {
	// flatten url.Values by dropping others but the first one
	values := make(map[string]string, len(q))
	for k, arr := range q {
		values[k] = arr[0]
	}
	return values
}

func marshalResponse(h http.Header, m proto.Message) ([]byte, error) {
	ct := h.Get("Content-Type")
	if strings.Contains(ct, headerProtobuf) {
		return proto.Marshal(m)
	} else if strings.Contains(ct, headerJSON) {
		return json.Marshal(m)
	}
	return nil, fmt.Errorf("on marshalBody: invalid content-type; type(%v)", ct)
}

func unmarshalBody(h http.Header, r io.Reader, m proto.Message) error {
	ct := h.Get("Content-Type")
	if strings.Contains(ct, headerProtobuf) {
		return unmarshalBodyProto(r, m)
	} else if strings.Contains(ct, headerJSON) {
		return unmarshalBodyJSON(r, m)
	}
	return fmt.Errorf("on unmarshalBody: invalid content-type; type(%v)", ct)
}

func unmarshalBodyProto(r io.Reader, m proto.Message) error {
	arr, err := ioutil.ReadAll(r)
	if err != nil {
		return fmt.Errorf("on unmarshalBodyProto: %v", err)
	}
	return puOption.Unmarshal(arr, m)
}

func unmarshalBodyJSON(r io.Reader, m proto.Message) error {
	if err := json.NewDecoder(r).Decode(m); err != nil {
		return fmt.Errorf("on unmarshalBodyJSON: %v", err)
	}
	return nil
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
