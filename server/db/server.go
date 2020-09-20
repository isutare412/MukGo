package db

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Server runs as Database server for MukGo service. Server should be created
// with NewServer function.
type Server struct {
	mqss *mq.Session

	db     *mongo.Database
	dbconn *mongo.Client
	dbctx  context.Context
}

var baseConfig = &mq.SessionConfig{
	Exchanges: map[string]mq.ExchangeConfig{
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
		dbctx: context.Background(),
	}

	// build option for MongoDB
	uri := fmt.Sprintf(
		"mongodb://%s:%s@%s:%d",
		cfg.MongoDB.User,
		cfg.MongoDB.Password,
		cfg.MongoDB.IP,
		cfg.MongoDB.Port,
	)
	option := options.Client().ApplyURI(uri)

	// connect to MongoDB
	client, err := mongo.Connect(s.dbctx, option)
	if err != nil {
		return nil, fmt.Errorf("on newDBConn: %v", err)
	}

	// check the connection
	err = client.Ping(s.dbctx, nil)
	if err != nil {
		return nil, fmt.Errorf("on newDBConn: %v", err)
	}
	s.dbconn = client
	console.Info("MongoDB connection established")

	// select database to use
	s.db = client.Database("mukgo")

	// establish rabbitmq session
	mqaddr := fmt.Sprintf("%s:%d", cfg.RabbitMQ.IP, cfg.RabbitMQ.Port)
	baseConfig.User = cfg.RabbitMQ.User
	baseConfig.Password = cfg.RabbitMQ.Password
	baseConfig.Addr = mqaddr
	session := mq.NewSession("db", baseConfig)

	// connect the session
	if err := session.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	s.mqss = session
	console.Info("session(%s) established between RabbitMQ", mqaddr)

	return s, nil
}

// Run start handling database requests.
func (s *Server) Run() error {
	err := s.mqss.Consume(server.MGDB, server.API2DB, s.handleDBRequest)
	if err != nil {
		return fmt.Errorf("on run: %v", err)
	}

	// wait forever
	<-make(chan struct{})

	return nil
}

func (s *Server) handleDBRequest(d *amqp.Delivery) (res bool, err error) {
	var response server.Packet = &server.PacketError{}
	defer func() {
		if pubErr := s.mqss.Reply(
			server.MGDB,
			d.ReplyTo,
			server.DB,
			d.CorrelationId,
			response,
		); pubErr != nil {
			res = false
			err = pubErr
		}
	}()

	_, packetType, err := mq.ParseHeader(d.Headers)
	if err != nil {
		return false, fmt.Errorf("on handleDBRequest: %v", err)
	}

	// parse packet
	switch packetType {
	case server.PTReview:
		var p server.PacketReview
		err = json.Unmarshal(d.Body, &p)
		if err != nil {
			break
		}
		response = s.handleReview(&p)

	default:
		err = fmt.Errorf("no parser for %v", packetType)
	}

	// packet handling failed
	if err != nil {
		return false, fmt.Errorf("on handleDBRequest: %v", err)
	}

	return true, nil
}

// handleReview insert new review.
func (s *Server) handleReview(p *server.PacketReview) server.Packet {
	collection := s.db.Collection("reviews")

	_, err := collection.InsertOne(s.dbctx, struct {
		UserID  int
		Score   int
		Comment string
	}{
		p.UserID,
		p.Score,
		p.Comment,
	})
	if err != nil {
		return &server.PacketError{Message: "failed to insert review"}
	}

	return &server.PacketAck{}
}
