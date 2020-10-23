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
		return &server.DAPacketError{ErrorType: server.ETUserExists}
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
			return &server.DAPacketError{ErrorType: server.ETNoSuchUser}
		default:
			console.Warning(
				"on handleUserGet: failed to get user; packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	console.Info("send user data; User(%v)", *user)
	return &server.DAPacketUser{
		UserID: user.UserID,
		Name:   user.Name,
		Exp:    user.Exp,
	}
}

func (s *Server) handleReviewsGet(p *server.ADPacketReviewsGet) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	// check restaurant data
	restaurant, err := queryRestaurantGet(ctx, s.db, p.RestID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleReviewsGet: cannot find restaurant;"+
					"packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETNoSuchRestaurant}
		default:
			console.Warning(
				"on handleReviewsGet: failed to get restaurant; packet(%v): %v",
				*p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	reviews, err := queryReviewsGet(ctx, s.db, p.RestID)
	if err != nil {
		console.Warning(
			"on handleReviewsGet: failed to get reviews; packet(%v): %v",
			*p, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	// find user name from user id.
	var idMap = make(map[string]string)
	for _, r := range reviews {
		// check if name already cached
		if _, ok := idMap[r.UserID]; ok {
			continue
		}

		user, err := queryUserGet(ctx, s.db, r.UserID)
		if err != nil {
			switch err {
			case mongo.ErrNoDocuments:
				// database integrity contraint broken
				console.Error(
					"on handleReviewsGet: cannot find user; uid(%v): %v",
					r.UserID, err)
				return &server.DAPacketError{ErrorType: server.ETInternal}
			default:
				console.Warning(
					"on handleReviewsGet: failed to get user; uid(%v): %v",
					r.UserID, err)
				return &server.DAPacketError{ErrorType: server.ETInternal}
			}
		}

		idMap[r.UserID] = user.Name
	}

	// copy review data into response packet
	resPacket := server.DAPacketReviews{
		Reviews: make([]*common.Review, 0, len(reviews)),
	}
	for _, r := range reviews {
		resPacket.Reviews = append(resPacket.Reviews,
			&common.Review{
				UserID:   r.UserID,
				UserName: idMap[r.UserID],
				Score:    r.Score,
				Comment:  r.Comment,
			},
		)
	}

	console.Info("send reviews; restaurant(%v) count(%v)",
		restaurant.Name, len(resPacket.Reviews))
	return &resPacket
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
			return &server.DAPacketError{ErrorType: server.ETNoSuchUser}
		default:
			console.Warning(
				"on handleReviewAdd: failed to get user; packet(%v): %v",
				*p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
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
			return &server.DAPacketError{ErrorType: server.ETNoSuchRestaurant}
		default:
			console.Warning(
				"on handleReviewAdd: failed to get restaurant; packet(%v): %v",
				*p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// add review data
	err = queryReviewAdd(ctx, s.db, p.UserID, p.RestID, p.Score, p.Comment)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed to insert review(%v): %v", *p, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	// add exp to user
	user.Exp += common.ReviewExp()
	err = queryUserUpdate(ctx, s.db, user)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed update user; User(%v): %v", *user, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	console.Info("insert review; UserID(%v), Score(%v)", p.UserID, p.Score)
	return &server.DAPacketUser{
		UserID: user.UserID,
		Name:   user.Name,
		Exp:    user.Exp,
	}
}

func (s *Server) handleRestaurantGet(
	p *server.ADPacketRestaurantGet,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	restaurant, err := queryRestaurantGet(ctx, s.db, p.RestID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleRestaurantGet: cannot find restaurant; "+
					"packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETNoSuchRestaurant}
		default:
			console.Warning(
				"on handleRestaurantGet: failed to get restaurant; "+
					"packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	console.Info("found restaurant; Name(%v)", *restaurant)
	return &server.DAPacketRestaurant{
		Restaurant: &common.Restaurant{
			ID:   restaurant.ID,
			Name: restaurant.Name,
			Coord: common.Coordinate{
				Latitude:  restaurant.Latitude,
				Longitude: restaurant.Longitude,
			},
		},
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
		return &server.DAPacketError{ErrorType: server.ETInternal}
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
			return &server.DAPacketError{ErrorType: server.ETNoSuchUser}
		default:
			console.Warning(
				"on handleRestaurantsGet: failed to get user; packet(%v)", *p)
			return &server.DAPacketError{ErrorType: server.ETInternal}
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
		return &server.DAPacketError{ErrorType: server.ETInternal}
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
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	console.Info("insert restaurants; count(%v)", len(p.Restaurants))
	return &server.DAPacketAck{}
}
