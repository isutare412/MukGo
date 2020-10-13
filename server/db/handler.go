package db

import (
	"context"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const queryTimeout = 5 * time.Second

func (s *Server) handleUserAdd(p *server.ADPacketUserAdd) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	coll := s.db.Collection(CNUser)
	_, err := coll.InsertOne(
		ctx,
		User{
			UserID: p.UserID,
			Name:   p.Name,
			Exp:    0,
		})
	if err != nil {
		console.Warning("failed to insert user(%v): %v", *p, err)
		return &server.DAPacketUserExist{UserID: p.UserID}
	}

	console.Info("insert user; UserID(%v), Name(%v)", p.UserID, p.Name)
	return &server.DAPacketAck{}
}

func (s *Server) handleUserGet(p *server.ADPacketUserGet) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	coll := s.db.Collection(CNUser)
	cursor := coll.FindOne(
		ctx,
		bson.M{
			"userid": p.UserID,
		})

	switch cursor.Err() {
	case nil:
		break // success
	case mongo.ErrNoDocuments:
		console.Warning("cannot find user; packet(%v)", *p)
		return &server.DAPacketNoSuchUser{UserID: p.UserID}
	default:
		console.Warning("failed to find user; packet(%v)", *p)
		return &server.DAPacketError{Message: "cannot find user"}
	}

	var found User
	if err := cursor.Decode(&found); err != nil {
		console.Warning("failed to decode user; packet(%v): %v", *p, err)
		return &server.DAPacketError{Message: "failed to decode user"}
	}

	console.Info("send user data; User(%v)", found)
	return &server.DAPacketUser{
		UserID: found.UserID,
		Name:   found.Name,
		Exp:    found.Exp,
	}
}

func (s *Server) handleReviewAdd(p *server.ADPacketReviewAdd) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	coll := s.db.Collection(CNReview)
	_, err := coll.InsertOne(
		ctx,
		Review{
			UserID:  p.UserID,
			Score:   p.Score,
			Comment: p.Comment,
		})
	if err != nil {
		console.Warning("failed to insert review(%v): %v", *p, err)
		return &server.DAPacketError{Message: "failed to insert review"}
	}

	console.Info("insert review; UserID(%v), Score(%v)", p.UserID, p.Score)
	return &server.DAPacketAck{}
}

func (s *Server) handleRestaurantAdd(p *server.ADPacketRestaurantAdd) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	coll := s.db.Collection(CNRestaurant)
	_, err := coll.InsertOne(
		ctx,
		Restaurant{
			Name:      p.Name,
			Latitude:  p.Coord.Latitude,
			Longitude: p.Coord.Longitude,
		})
	if err != nil {
		console.Warning("failed to insert restaurant(%v): %v", *p, err)
		return &server.DAPacketError{Message: "failed to insert restaurant"}
	}

	console.Info("insert restaurant; Name(%v)", p.Name)
	return &server.DAPacketAck{}
}
