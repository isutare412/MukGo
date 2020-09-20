package db

import (
	"context"
	"fmt"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
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
	var server = &Server{
		dbctx: context.Background(),
	}

	// establish rabbitmq session
	mqaddr := fmt.Sprintf("%s:%d", cfg.RabbitMQ.IP, cfg.RabbitMQ.Port)
	baseConfig.User = cfg.RabbitMQ.User
	baseConfig.Password = cfg.RabbitMQ.Password
	baseConfig.Addr = mqaddr
	session := mq.NewSession("db", baseConfig)

	// connection the session
	if err := session.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	server.mqss = session
	console.Info("session(%s) established between RabbitMQ", mqaddr)

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
	client, err := mongo.Connect(server.dbctx, option)
	if err != nil {
		return nil, fmt.Errorf("on newDBConn: %v", err)
	}

	// check the connection
	err = client.Ping(server.dbctx, nil)
	if err != nil {
		return nil, fmt.Errorf("on newDBConn: %v", err)
	}
	server.dbconn = client
	console.Info("MongoDB connection established")

	// select database to use
	server.db = client.Database("mukgo")

	return server, nil
}

func (s *Server) TestQuery() {
	collection := s.db.Collection("trainers")
	collection.InsertOne(s.dbctx, struct {
		Name string
		Age  int
	}{"Redshore", 18})
}
