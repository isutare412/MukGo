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
	"go.mongodb.org/mongo-driver/bson"
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
	console.Info("connect to MongoDB...")
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
	console.Info("connect to RabbitMQ...")
	if err := session.TryConnect(40, 3000*time.Millisecond); err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	s.mqss = session
	console.Info("session(%s) established between RabbitMQ", mqaddr)

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
				server.DB,
				&packet,
			); err != nil {
				return false
			}
			return true
		},
	)

	return s, nil
}

// InitDB creates database, collections, indexex.
func (s *Server) InitDB() error {
	collUser := s.db.Collection(CNUser)
	_, err := collUser.Indexes().CreateOne(
		s.dbctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "user_id", Value: 1}, // index in ascending order
			},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		return fmt.Errorf("on InitDB: %v", err)
	}

	collLike := s.db.Collection(CNLike)
	_, err = collLike.Indexes().CreateOne(
		s.dbctx,
		mongo.IndexModel{
			Keys: bson.D{
				{Key: "review_id", Value: 1},      // index in ascending order
				{Key: "liking_user_id", Value: 1}, // index in ascending order
			},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		return fmt.Errorf("on InitDB: %v", err)
	}

	return nil
}

// Run start handling database requests.
func (s *Server) Run() error {
	err := s.mqss.Consume(server.MGDB, server.API2DB, s.handleDBRequest, 4)
	if err != nil {
		return fmt.Errorf("on run: %v", err)
	}

	// wait forever
	<-make(chan struct{})

	return nil
}

func (s *Server) handleDBRequest(d *amqp.Delivery) (res bool, err error) {
	var response server.Packet = &server.DAPacketError{}
	defer func() {
		// reply RPC with response packet
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
	response, err = s.handlePacket(packetType, d.Body)
	if err != nil {
		return false, fmt.Errorf("on handleDBRequest: %v", err)
	}

	return true, nil
}

func (s *Server) handlePacket(
	pt server.PacketType, ser []byte,
) (response server.Packet, err error) {
	// parse packet
	switch pt {
	case server.PTADUserAdd:
		var p server.ADPacketUserAdd
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleUserAdd(&p)

	case server.PTADUserGet:
		var p server.ADPacketUserGet
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleUserGet(&p)

	case server.PTADReviewGet:
		var p server.ADPacketReviewGet
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleReviewGet(&p)

	case server.PTADReviewsGet:
		var p server.ADPacketReviewsGet
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleReviewsGet(&p)

	case server.PTADReviewAdd:
		var p server.ADPacketReviewAdd
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleReviewAdd(&p)

	case server.PTADReviewDel:
		var p server.ADPacketReviewDel
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleReviewDel(&p)

	case server.PTADRestaurantGet:
		var p server.ADPacketRestaurantGet
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleRestaurantGet(&p)

	case server.PTADRestaurantAdd:
		var p server.ADPacketRestaurantAdd
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleRestaurantAdd(&p)

	case server.PTADRestaurantsGet:
		var p server.ADPacketRestaurantsGet
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleRestaurantsGet(&p)

	case server.PTADRestaurantsAdd:
		var p server.ADPacketRestaurantsAdd
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleRestaurantsAdd(&p)

	case server.PTADRankingGet:
		var p server.ADPacketRankingGet
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleRankingGet(&p)

	case server.PTADLikeAdd:
		var p server.ADPacketLikeAdd
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleLikeAdd(&p)

	case server.PTADLikeDel:
		var p server.ADPacketLikeDel
		err = json.Unmarshal(ser, &p)
		if err != nil {
			err = fmt.Errorf("on handlePacket: %v", err)
			break
		}
		response = s.handleLikeDel(&p)

	default:
		err = fmt.Errorf("on handlePacket: no parser for %v", pt)
	}

	return
}
