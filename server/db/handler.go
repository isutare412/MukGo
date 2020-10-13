package db

import (
	"context"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
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
		return &server.DAPacketError{Message: "failed to insert user"}
	}

	console.Info("insert user; UserID(%d), Name(%s)", p.UserID, p.Name)
	return &server.DAPacketAck{}
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

	console.Info("insert review; UserID(%d), Score(%d)", p.UserID, p.Score)
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

	console.Info("insert restaurant; Name(%s)", p.Name)
	return &server.DAPacketAck{}
}
