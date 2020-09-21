package log

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
	"github.com/isutare412/dailyrotate"
	"github.com/streadway/amqp"
)

// Server runs as log server for MukGo service. Server should be created with
// NewServer function.
type Server struct {
	mqss    *mq.Session
	loggers map[string]*logger
}

var baseConfig = &mq.SessionConfig{
	Exchanges: map[string]mq.ExchangeConfig{
		server.MGLogs: {
			Name: server.MGLogs,
			Type: "fanout",
			Queues: map[string]mq.QueueConfig{
				server.Log: {
					Name:       "", // will use generated name
					Durable:    false,
					AutoDelete: true,
				},
			},
		},
	},
}

// NewServer creates log server struct safely.
func NewServer(cfg *ServerConfig) (*Server, error) {
	var server = &Server{
		loggers: map[string]*logger{
			server.API: {
				sender: server.API,
				tag:    "API",
				dir:    "logs/api",
			},
			server.DB: {
				sender: server.DB,
				tag:    "DB",
				dir:    "logs/db",
			},
		},
	}

	// establish rabbitmq session
	mqaddr := fmt.Sprintf("%s:%d", cfg.RabbitMQ.IP, cfg.RabbitMQ.Port)
	baseConfig.User = cfg.RabbitMQ.User
	baseConfig.Password = cfg.RabbitMQ.Password
	baseConfig.Addr = mqaddr
	mqSession := mq.NewSession("api", baseConfig)

	// connection the session
	console.Info("connect to RabbitMQ...")
	if err := mqSession.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	server.mqss = mqSession
	console.Info("session(%q) established between RabbitMQ", mqaddr)

	// create direcotries for logging
	for _, l := range server.loggers {
		if err := os.MkdirAll(l.dir, 0777); err != nil {
			return nil, fmt.Errorf("on NewServer: %v", err)
		}
	}

	// load location for log setup
	loc, err := time.LoadLocation("Asia/Seoul")
	if err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}

	// open log files
	for _, l := range server.loggers {
		file, err := dailyrotate.NewFile(
			fmt.Sprintf("%s/%s", l.dir, "2006-01-02-150405.log"), loc, nil)
		if err != nil {
			return nil, fmt.Errorf("on NewServer: %v", err)
		}
		l.file = file
		console.Info("log file(%q) created", file.Path())
	}

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
	sender, _, err := mq.ParseHeader(d.Headers)
	if err != nil {
		return false, fmt.Errorf("on handleLog: %v", err)
	}

	var packet server.PacketLog
	if err := json.Unmarshal(d.Body, &packet); err != nil {
		return false, fmt.Errorf("failed to unmarshall packet")
	}

	// select proper logger from sender
	logger, ok := s.loggers[sender]
	if !ok {
		return false, fmt.Errorf("cannot handle unidentified sender(%s)", sender)
	}

	// now leave log
	if err := logger.log(packet.Timestamp, packet.Msg); err != nil {
		return false, fmt.Errorf("on handleLog: %v", err)
	}

	return true, nil
}
