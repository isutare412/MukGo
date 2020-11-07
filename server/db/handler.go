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

	err := queryUserAdd(ctx, s.db, p.UserID, p.Name, 0, 0, 0)
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

	// find restaurant counts if heavy request
	var rtCounts map[int32]int32

	// queries for heavy request
	if p.HeavyRequest {
		rtCounts = make(map[int32]int32)

		// find review by user
		reviews, err := queryReviewsGetByUser(ctx, s.db, p.UserID)
		if err != nil {
			console.Warning(
				"on handleUserGet: failed to get reviews; packet(%v): %v",
				*p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}

		for _, review := range reviews {
			// find restaurant data
			restaurant, err := queryRestaurantGet(ctx, s.db, review.RestID)
			if err != nil {
				switch err {
				case mongo.ErrNoDocuments:
					console.Warning(
						"on handleUserGet: cannot find restaurant;"+
							" packet(%v): %v", *p, err)
					return &server.DAPacketError{ErrorType: server.ETInternal}
				default:
					console.Warning(
						"on handleUserGet: failed to get restaurant; packet(%v): %v",
						*p, err)
					return &server.DAPacketError{ErrorType: server.ETInternal}
				}
			}

			// count group by restaurant type
			if restaurant.Type != 0 {
				rtCounts[restaurant.Type]++
			}
		}

		console.Info("send user heavy data; User(%v)", user.Name)
	}

	// console.Info("send user data; User(%v)", *user)
	return &server.DAPacketUser{
		User: &common.User{
			UserID:      user.UserID,
			Name:        user.Name,
			Exp:         user.Exp,
			ReviewCount: user.ReviewCount,
			LikeCount:   user.LikeCount,
			RTCounts:    rtCounts,
		},
	}
}

func (s *Server) handleReviewGet(p *server.ADPacketReviewGet) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	// find review
	review, err := queryReviewGet(ctx, s.db, p.ReviewID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			// database integrity contraint broken
			console.Error(
				"on handleReviewGet: cannot find review; rid(%v): %v",
				p.ReviewID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		default:
			console.Warning(
				"on handleReviewGet: failed to get review; rid(%v): %v",
				p.ReviewID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// get liked user
	likedUser, err := queryUserGet(ctx, s.db, review.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleReviewGet: cannot find liked user; uid(%v): %v",
				review.UserID, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		default:
			console.Warning(
				"on handleReviewGet: failed to get liked user; uid(%v): %v",
				review.UserID, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// count likes of review
	var likeCount int32
	likes, err := queryLikesGetByReview(ctx, s.db, review.ID)
	if err != nil {
		console.Warning(
			"on handleReviewGet: failed to get likes of review; rid(%v): %v",
			review.ID.Hex(), err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}
	likeCount = int32(len(likes))

	var likedByRequester bool
	for _, l := range likes {
		if l.LikingUserID == p.UserID {
			likedByRequester = true
		}
	}

	// send review data
	console.Info("send review; rid(%v)", p.ReviewID.Hex())
	return &server.DAPacketReview{
		Review: &common.Review{
			ID:        review.ID,
			UserID:    review.UserID,
			UserName:  likedUser.Name,
			UserExp:   likedUser.Exp,
			Score:     review.Score,
			Comment:   review.Comment,
			Menus:     review.Menus,
			Wait:      review.Wait,
			NumPeople: review.NumPeople,
			Timestamp: review.Timestamp,
			LikeCount: likeCount,
			LikedByMe: likedByRequester,
		},
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

	reviews, err := queryReviewsGetByRestaurant(ctx, s.db, p.RestID)
	if err != nil {
		console.Warning(
			"on handleReviewsGet: failed to get reviews; packet(%v): %v",
			*p, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	// collect review related data
	var userMap = make(map[string]*User)
	var likeCounts = make(map[*Review]int32)
	var likeByUser = make(map[*Review]bool)
	for _, r := range reviews {
		// count likes of each review
		likes, err := queryLikesGetByReview(ctx, s.db, r.ID)
		if err != nil {
			console.Warning(
				"on handleReviewsGet: failed to get likes of liked user; "+
					"uid(%v): %v",
				r.ID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
		likeCounts[r] += int32(len(likes))

		// check review is liked by requester
		for _, l := range likes {
			if l.LikingUserID == p.UserID {
				likeByUser[r] = true
			}
		}

		// check if name already cached
		if _, ok := userMap[r.UserID]; ok {
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
		userMap[r.UserID] = user
	}

	// copy review data into response packet
	resPacket := server.DAPacketReviews{
		Reviews: make([]*common.Review, 0, len(reviews)),
	}
	for _, r := range reviews {
		resPacket.Reviews = append(resPacket.Reviews,
			&common.Review{
				ID:        r.ID,
				UserID:    r.UserID,
				UserName:  userMap[r.UserID].Name,
				UserExp:   userMap[r.UserID].Exp,
				Score:     r.Score,
				Comment:   r.Comment,
				Menus:     r.Menus,
				Wait:      r.Wait,
				NumPeople: r.NumPeople,
				Timestamp: r.Timestamp,
				LikeCount: likeCounts[r],
				LikedByMe: likeByUser[r],
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

	// count review by user
	reviews, err := queryReviewsGetByUser(ctx, s.db, p.UserID)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed to get reviews; packet(%v): %v",
			*p, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}
	reviewCount := int32(len(reviews))

	// add review data
	err = queryReviewAdd(ctx, s.db, p.UserID, p.RestID, p.Score, p.Comment,
		p.Menus, p.Wait, p.NumPeople, p.Timestamp)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed to insert review(%v): %v", *p, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	// update user user data
	user.Exp += common.ReviewExp()
	user.ReviewCount = reviewCount + 1
	err = queryUserUpdate(ctx, s.db, user)
	if err != nil {
		console.Warning(
			"on handleReviewAdd: failed update user; User(%v): %v", *user, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	console.Info("insert review; UserID(%v), Score(%v)", p.UserID, p.Score)
	return &server.DAPacketUser{
		User: &common.User{
			UserID:      user.UserID,
			Name:        user.Name,
			Exp:         user.Exp,
			ReviewCount: user.ReviewCount,
			LikeCount:   user.LikeCount,
		},
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
			RestaurantType: restaurant.Type,
		},
	}
}

func (s *Server) handleRestaurantAdd(
	p *server.ADPacketRestaurantAdd,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	err := queryRestaurantAdd(
		ctx, s.db, p.Name, p.Coord.Latitude, p.Coord.Longitude,
		p.RestaurantType)
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
				RestaurantType: r.Type,
			})
	}

	// console.Info("found restaurants; count(%v)", len(resPacket.Restaurants))
	return &resPacket
}

func (s *Server) handleRestaurantsAdd(
	p *server.ADPacketRestaurantsAdd,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	for _, r := range p.Restaurants {
		err := queryRestaurantAdd(
			ctx, s.db, r.Name, r.Coord.Latitude, r.Coord.Longitude,
			r.RestaurantType)

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

func (s *Server) handleRankingGet(
	p *server.ADPacketRankingGet,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	users, err := queryUserRankingGet(ctx, s.db, 10)
	if err != nil {
		console.Warning("on handleRankingGet: failed to get ranking: %v", err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	pusers := make([]*common.User, 0, len(users))
	for _, u := range users {
		pusers = append(pusers,
			&common.User{
				UserID:      u.UserID,
				Name:        u.Name,
				Exp:         u.Exp,
				ReviewCount: u.ReviewCount,
				LikeCount:   u.LikeCount,
			},
		)
	}

	console.Info("send rankings; count(%v)", len(pusers))
	return &server.DAPacketUsers{Users: pusers}
}

func (s *Server) handleLikeAdd(
	p *server.ADPacketLikeAdd,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	// check request user exists
	_, err := queryUserGet(ctx, s.db, p.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleLikeAdd: cannot find user; packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETNoSuchUser}
		default:
			console.Warning(
				"on handleLikeAdd: failed to get user; packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// check like requested review exists
	review, err := queryReviewGet(ctx, s.db, p.ReviewID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			// database integrity contraint broken
			console.Error(
				"on handleLikeAdd: cannot find review; rid(%v): %v",
				p.ReviewID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		default:
			console.Warning(
				"on handleLikeAdd: failed to get review; rid(%v): %v",
				p.ReviewID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// get liked user
	likedUser, err := queryUserGet(ctx, s.db, review.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleLikeAdd: cannot find liked user; uid(%v): %v",
				review.UserID, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		default:
			console.Warning(
				"on handleLikeAdd: failed to get liked user; uid(%v): %v",
				review.UserID, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// count likes of liked user
	likes, err := queryLikesGetByLikedUser(ctx, s.db, likedUser.UserID)
	if err != nil {
		console.Warning(
			"on handleLikeAdd: failed to get likes of liked user; uid(%v): %v",
			likedUser.UserID, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}
	likeCount := int32(len(likes))

	// add like data
	err = queryLikeAdd(ctx, s.db, p.UserID, likedUser.UserID, p.ReviewID)
	if err != nil {
		console.Warning(
			"on handleLikeAdd: failed to insert like(%v): %v", *p, err)
		return &server.DAPacketError{ErrorType: server.ETLikeExists}
	}

	// update user user data
	likedUser.Exp += common.LikeExp()
	likedUser.LikeCount = likeCount + 1
	err = queryUserUpdate(ctx, s.db, likedUser)
	if err != nil {
		console.Warning(
			"on handleLikeAdd: failed update liked user; User(%v): %v",
			*likedUser, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	// count likes of review
	likeCount = 0
	likes, err = queryLikesGetByReview(ctx, s.db, review.ID)
	if err != nil {
		console.Warning(
			"on handleLikeAdd: failed to get likes of review; "+
				"rid(%v): %v",
			review.ID.Hex(), err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}
	likeCount = int32(len(likes))

	console.Info("saved like; uid(%v) rid(%v)", p.UserID, p.ReviewID.Hex())
	return &server.DAPacketReview{
		Review: &common.Review{
			ID:        review.ID,
			UserID:    review.UserID,
			UserName:  likedUser.Name,
			UserExp:   likedUser.Exp,
			Score:     review.Score,
			Comment:   review.Comment,
			Menus:     review.Menus,
			Wait:      review.Wait,
			NumPeople: review.NumPeople,
			Timestamp: review.Timestamp,
			LikeCount: likeCount,
			LikedByMe: true,
		},
	}
}

func (s *Server) handleLikeDel(
	p *server.ADPacketLikeDel,
) server.Packet {
	ctx, cancel := context.WithTimeout(s.dbctx, queryTimeout)
	defer cancel()

	// check request user exists
	_, err := queryUserGet(ctx, s.db, p.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleLikeDel: cannot find user; packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETNoSuchUser}
		default:
			console.Warning(
				"on handleLikeDel: failed to get user; packet(%v): %v", *p, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// check like requested review exists
	review, err := queryReviewGet(ctx, s.db, p.ReviewID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			// database integrity contraint broken
			console.Error(
				"on handleLikeDel: cannot find review; rid(%v): %v",
				p.ReviewID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		default:
			console.Warning(
				"on handleLikeDel: failed to get review; rid(%v): %v",
				p.ReviewID.Hex(), err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// get liked user
	likedUser, err := queryUserGet(ctx, s.db, review.UserID)
	if err != nil {
		switch err {
		case mongo.ErrNoDocuments:
			console.Warning(
				"on handleLikeDel: cannot find liked user; uid(%v): %v",
				review.UserID, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		default:
			console.Warning(
				"on handleLikeDel: failed to get liked user; uid(%v): %v",
				review.UserID, err)
			return &server.DAPacketError{ErrorType: server.ETInternal}
		}
	}

	// count likes of liked user
	likes, err := queryLikesGetByLikedUser(ctx, s.db, likedUser.UserID)
	if err != nil {
		console.Warning(
			"on handleLikeDel: failed to get likes of liked user; uid(%v): %v",
			likedUser.UserID, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}
	likeCount := int32(len(likes))

	// delete like data
	err = queryLikeDel(ctx, s.db, p.UserID, p.ReviewID)
	if err != nil {
		console.Warning(
			"on handleLikeDel: failed to delete like(%v): %v", *p, err)
		return &server.DAPacketError{ErrorType: server.ETLikeExists}
	}

	// update user user data
	likedUser.Exp -= common.LikeExp()
	if likedUser.Exp < 0 {
		likedUser.Exp = 0
	}
	likedUser.LikeCount = likeCount - 1
	err = queryUserUpdate(ctx, s.db, likedUser)
	if err != nil {
		console.Warning(
			"on handleLikeDel: failed update liked user; User(%v): %v",
			*likedUser, err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}

	// count likes of review
	likeCount = 0
	likes, err = queryLikesGetByReview(ctx, s.db, review.ID)
	if err != nil {
		console.Warning(
			"on handleLikeDel: failed to get likes of review; "+
				"rid(%v): %v",
			review.ID.Hex(), err)
		return &server.DAPacketError{ErrorType: server.ETInternal}
	}
	likeCount = int32(len(likes))

	console.Info("saved like; uid(%v) rid(%v)", p.UserID, p.ReviewID.Hex())
	return &server.DAPacketReview{
		Review: &common.Review{
			ID:        review.ID,
			UserID:    review.UserID,
			UserName:  likedUser.Name,
			UserExp:   likedUser.Exp,
			Score:     review.Score,
			Comment:   review.Comment,
			Menus:     review.Menus,
			Wait:      review.Wait,
			NumPeople: review.NumPeople,
			Timestamp: review.Timestamp,
			LikeCount: likeCount,
			LikedByMe: false,
		},
	}
}
