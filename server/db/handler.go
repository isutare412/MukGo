package db

import (
	"context"
	"time"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/common"
	"github.com/isutare412/MukGo/server/console"
	"go.mongodb.org/mongo-driver/mongo"
)

const queryTimeout = 5 * time.Second

func (s *Server) handleUserAdd(p *server.ADPacketUserAdd) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	err := queryUserAdd(ctx, s.db, p.UserID, p.Name, 0)
	if err != nil {
		console.Warning(
			"on handleUserAdd: failed to insert user(%v): %v", *p, err)
		return &server.DAPacketUserExist{UserID: p.UserID}
	}

	console.Info("insert user; UserID(%v), Name(%v)", p.UserID, p.Name)
	return &server.DAPacketAck{}
}

func (s *Server) handleUserGet(p *server.ADPacketUserGet) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	user, err := queryUserGet(ctx, s.db, p.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleUserGet: cannot find user; packet(%v): %v", *p, err)
			return &server.DAPacketNoSuchUser{UserID: p.UserID}
		default:
			console.Warning(
				"on handleUserGet: failed to get user; packet(%v): %v", *p, err)
			return &server.DAPacketError{Message: "failed to get user"}
		}
	}

	console.Info("send user data; User(%v)", *user)
	return &server.DAPacketUser{
		UserID: user.UserID,
		Name:   user.Name,
		Exp:    user.Exp,
	}
}

func (s *Server) handleReviewAdd(p *server.ADPacketReviewAdd) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	// check user data
	user, err := queryUserGet(ctx, s.db, p.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleReviewAdd: cannot find user; packet(%v): %v", *p, err)
			return &server.DAPacketNoSuchUser{UserID: p.UserID}
		default:
			console.Warning(
				"on handleReviewAdd: failed to get user; packet(%v): %v",
				*p, err)
			return &server.DAPacketError{Message: "failed to get user"}
		}
	}

	// check restaurant data
	_, err = queryRestaurantGet(ctx, s.db, p.RestID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleReviewAdd: cannot find restaurant;"+
					"packet(%v): %v", *p, err)
			return &server.DAPacketNoSuchRestaurant{ID: p.RestID}
		default:
			console.Warning(
				"on handleReviewAdd: failed to get restaurant; packet(%v): %v",
				*p, err)
			return &server.DAPacketError{Message: "failed to get restaurant"}
		}
	}

	// add review data
	err = queryReviewAdd(ctx, s.db, p.UserID, p.RestID, p.Score, p.Comment)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed to insert review(%v): %v", *p, err)
		return &server.DAPacketError{Message: "failed to insert review"}
	}

	// add exp to user
	user.Exp += common.ReviewExp()
	err = queryUserUpdate(ctx, s.db, user)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed update user; User(%v): %v", *user, err)
		return &server.DAPacketError{Message: "failed to update user exp"}
	}

	console.Info("insert review; UserID(%v), Score(%v)", p.UserID, p.Score)
	return &server.DAPacketUser{
		UserID: user.UserID,
		Name:   user.Name,
		Exp:    user.Exp,
	}
}

func (s *Server) handleRestaurantAdd(
	p *server.ADPacketRestaurantAdd,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	err := queryRestaurantAdd(
		ctx, s.db, p.Name, p.Coord.Latitude, p.Coord.Longitude)
	if err != nil {
		console.Warning(
			"on handleRestaurantAdd: failed to insert restaurant(%v): %v",
			*p, err)
		return &server.DAPacketError{Message: "failed to insert restaurant"}
	}

	console.Info("insert restaurant; Name(%v)", p.Name)
	return &server.DAPacketAck{}
}

func (s *Server) handleRestaurantsGet(
	p *server.ADPacketRestaurantsGet,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	user, err := queryUserGet(ctx, s.db, p.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleRestaurantsGet: cannot find user; packet(%v)", *p)
			return &server.DAPacketNoSuchUser{UserID: p.UserID}
		default:
			console.Warning(
				"on handleRestaurantsGet: failed to get user; packet(%v)", *p)
			return &server.DAPacketError{Message: "failed to get user"}
		}
	}

	// get restaurants within sight range
	sight := common.Exp2Sight(user.Exp)
	northWest, southEast := p.Coord.RangeSquare(sight)
	restaurants, err := queryRestaurantsGet(
		ctx, s.db,
		southEast.Latitude, northWest.Latitude,
		northWest.Longitude, southEast.Longitude,
	)
	if err != nil {
		console.Warning(
			"on handleRestaurantsGet: failed to find restaurants; "+
				"Coord(%v): %v", p.Coord, err)
		return &server.DAPacketError{Message: "failed to find restaurants"}
	}

	// copy restaurants data
	resPacket := server.DAPacketRestaurants{
		Restaurants: make([]*common.Restaurant, 0, len(restaurants)),
	}
	for _, r := range restaurants {
		resPacket.Restaurants = append(resPacket.Restaurants,
			&common.Restaurant{
				ID:   r.ID,
				Name: r.Name,
				Coord: common.Coordinate{
					Latitude:  r.Latitude,
					Longitude: r.Longitude,
				},
			})
	}

	console.Info("found restaurants; count(%v)", len(resPacket.Restaurants))
	return &resPacket
}

func (s *Server) handleRestaurantsAdd(
	p *server.ADPacketRestaurantsAdd,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	for _, r := range p.Restaurants {
		err := queryRestaurantAdd(
			ctx, s.db, r.Name, r.Coord.Latitude, r.Coord.Longitude)

		if err != nil {
			console.Warning(
				"on handleRestaurantsAdd: failed to insert restaurant(%v): %v",
				*r, err)
			return &server.DAPacketError{Message: "failed to insert restaurant"}
		}
	}

	console.Info("insert restaurants; count(%v)", len(p.Restaurants))
	return &server.DAPacketAck{}
}
