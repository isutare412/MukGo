package log

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
	"github.com/streadway/amqp"
)

// Server runs as log server for MukGo service. Server should be created with
// NewServer function.
type Server struct {
	mqss *mq.Session
}

var baseConfig = &mq.SessionConfig{
	Exchanges: map[string]mq.ExchangeConfig{
		server.MGLogs: {
			Name: server.MGLogs,
			Type: "direct",
			Queues: map[string]mq.QueueConfig{
				server.Log: {
					Name:     server.Log,
					RouteKey: server.Log,
				},
			},
		},
	},
}

// NewServer creates log server struct safely.
func NewServer(mqid, mqpw, mqaddr string) (*Server, error) {
	var server = &Server{}

	// establish rabbitmq session
	baseConfig.User = mqid
	baseConfig.Password = mqpw
	baseConfig.Addr = mqaddr
	mqSession := mq.NewSession("api", baseConfig)

	// connection the session
	if err := mqSession.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	server.mqss = mqSession
	console.Info("session(%q) established between RabbitMQ", mqaddr)

	return server, nil
}

// Run start handling logs.
func (s *Server) Run() error {
	err := s.mqss.Consume(server.MGLogs, server.Log, s.handleLog)
	if err != nil {
		return fmt.Errorf("on HandleLogs: %v", err)
	}

	// wait forever
	<-make(chan struct{})

	return nil
}

func (s *Server) handleLog(d *amqp.Delivery) (bool, error) {
	header := d.Headers
	if header == nil {
		return false, fmt.Errorf("header does not exists in delievery")
	}

	sender, ok := header[server.Sender].(string)
	if !ok {
		return false, fmt.Errorf("sender does not exists in header")
	}

	var packet server.PacketLog
	if err := json.Unmarshal(d.Body, &packet); err != nil {
		return false, fmt.Errorf("failed to unmarshall packet")
	}

	switch sender {
	case server.API:
		s.handleAPILog(&packet)
	default:
		return false, fmt.Errorf(
			"cannot handle unidentified sender(%s)", sender)
	}

	return true, nil
}

func (s *Server) handleAPILog(packet *server.PacketLog) {
	console.Info(packet.Msg)
}
