package db

import (
	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
)

func (s *Server) handleUserAdd(p *server.PacketUserAdd) server.Packet {
	coll := s.db.Collection(CNUser)

	_, err := coll.InsertOne(
		s.dbctx,
		User{
			UserID: p.UserID,
			Name:   p.Name,
		})
	if err != nil {
		console.Warning("failed to insert user(%v): %v", *p, err)
		return &server.PacketError{Message: "failed to insert user"}
	}

	console.Info("insert user; UserID(%d), Name(%s)", p.UserID, p.Name)
	return &server.PacketAck{}
}

func (s *Server) handleReviewAdd(p *server.PacketReviewAdd) server.Packet {
	coll := s.db.Collection(CNReview)

	_, err := coll.InsertOne(
		s.dbctx,
		Review{
			UserID:  p.UserID,
			Score:   p.Score,
			Comment: p.Comment,
		})
	if err != nil {
		console.Warning("failed to insert review(%v): %v", *p, err)
		return &server.PacketError{Message: "failed to insert review"}
	}

	console.Info("insert review; UserID(%d), Score(%d)", p.UserID, p.Score)
	return &server.PacketAck{}
}

func (s *Server) handleRestaurantAdd(p *server.PacketRestaurantAdd) server.Packet {
	coll := s.db.Collection(CNRestaurant)

	_, err := coll.InsertOne(
		s.dbctx,
		Restaurant{
			Name:      p.Name,
			Latitude:  p.Latitude,
			Longitude: p.Longitude,
			Altitude:  p.Altitude,
		})
	if err != nil {
		console.Warning("failed to insert restaurant(%v): %v", *p, err)
		return &server.PacketError{Message: "failed to insert restaurant"}
	}

	console.Info("insert restaurant; Name(%s)", p.Name)
	return &server.PacketAck{}
}
