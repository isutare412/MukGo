package api

import (
	"net/http"
	"strconv"

	"github.com/isutare412/MukGo/server"
	pb "github.com/isutare412/MukGo/server/api/proto"
	"github.com/isutare412/MukGo/server/common"
	"github.com/isutare412/MukGo/server/console"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleUserGet(w, r)
	case "POST":
		s.handleUserPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, pb.Code_METHOD_NOT_ALLOWED)
	}
}

func (s *Server) handleUserGet(w http.ResponseWriter, r *http.Request) {
	uid, _, err := s.authenticate(r.Header)
	if err != nil {
		console.Warning("on handleUserGet: %v", err)
		httpError(w, http.StatusBadRequest, pb.Code_AUTH_FAILED)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketUserGet{
		UserID: uid,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleUserGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleUserGet: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleUserGet: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	case server.ETNoSuchUser:
		console.Warning("on handleUserGet: user not exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest, pb.Code_USER_NOT_EXISTS)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketUser)
	if !ok {
		console.Warning("on handleUserGet: unexpected packet")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// calculate level
	level, levExp, curExp, ratio := common.Exp2Level(packet.Exp)
	sightRadius := common.Level2Sight(level)

	// serialize user data
	ser, err := marshalResponse(r.Header,
		&pb.User{
			Id:          packet.UserID,
			Name:        packet.Name,
			Level:       level,
			TotalExp:    packet.Exp,
			LevelExp:    levExp,
			CurExp:      curExp,
			ExpRatio:    ratio,
			SightRadius: sightRadius,
			ReviewCount: packet.ReviewCount,
			LikeCount:   packet.LikeCount,
		})
	if err != nil {
		console.Warning("on handleUserGet: failed to marshal user data")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	baseHeader(w.Header())
	w.Write(ser)
	console.Info("sent user data; User(%v)", *packet)
}

func (s *Server) handleUserPost(w http.ResponseWriter, r *http.Request) {
	uid, name, err := s.authenticate(r.Header)
	if err != nil {
		console.Warning("on handleUserPost: %v", err)
		httpError(w, http.StatusBadRequest, pb.Code_AUTH_FAILED)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketUserAdd{
		UserID: uid,
		Name:   name,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleUserPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleUserPost: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleUserPost: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	case server.ETUserExists:
		console.Warning("on handleUserPost: user already exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest, pb.Code_USER_EXISTS)
		return
	}

	console.Info("new user created(%v)", name)
}

func (s *Server) handleReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.handleReviewPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, pb.Code_METHOD_NOT_ALLOWED)
	}
}

func (s *Server) handleReviewPost(w http.ResponseWriter, r *http.Request) {
	uid, _, err := s.authenticate(r.Header)
	if err != nil {
		console.Warning("on handleReviewPost: %v", err)
		httpError(w, http.StatusBadRequest, pb.Code_AUTH_FAILED)
		return
	}

	// parse request from client
	var userReq pb.ReviewPost
	err = unmarshalBody(r.Header, r.Body, &userReq)
	if err != nil || userReq.Review == nil {
		console.Warning("on handleReviewPost: failed to decode request")
		httpError(w, http.StatusBadRequest, pb.Code_PROTOCOL_MISMATCH)
		return
	}

	restID, err := primitive.ObjectIDFromHex(userReq.RestaurantId)
	if err != nil {
		console.Warning("on handleReviewPost: invalid restaurantd id; "+
			"id(%v): %v", userReq.RestaurantId, err)
		httpError(w, http.StatusBadRequest, pb.Code_INVALID_DATA)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketReviewAdd{
		UserID:    uid,
		RestID:    restID,
		Score:     userReq.Review.Score,
		Comment:   userReq.Review.Comment,
		Menus:     userReq.Review.Menus,
		Wait:      userReq.Review.Wait,
		NumPeople: userReq.Review.NumPeople,
		Timestamp: userReq.Review.Timestamp,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleReviewPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleReviewPost: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleReviewPost: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	case server.ETNoSuchUser:
		console.Warning("on handleReviewPost: user not exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest, pb.Code_USER_NOT_EXISTS)
		return
	case server.ETNoSuchRestaurant:
		console.Warning(
			"on handleReviewPost: restaurant not exists; RestID(%v)",
			dbReq.RestID)
		httpError(w, http.StatusBadRequest, pb.Code_RESTAURANT_NOT_EXISTS)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketUser)
	if !ok {
		console.Warning("on handleReviewPost: unexpected packet")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// calculate level
	level, levExp, curExp, ratio := common.Exp2Level(packet.Exp)
	sightRadius := common.Level2Sight(level)

	// serialize user data
	ser, err := marshalResponse(r.Header,
		&pb.User{
			Id:          packet.UserID,
			Name:        packet.Name,
			TotalExp:    packet.Exp,
			LevelExp:    levExp,
			CurExp:      curExp,
			ExpRatio:    ratio,
			SightRadius: sightRadius,
			ReviewCount: packet.ReviewCount,
			LikeCount:   packet.LikeCount,
		})
	if err != nil {
		console.Warning("on handleReviewPost: failed to marshal user data")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// send updated user data
	baseHeader(w.Header())
	w.Write(ser)
	console.Info("new review from user; User(%v)", *packet)
}

func (s *Server) handleReviews(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleReviewsGet(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, pb.Code_METHOD_NOT_ALLOWED)
	}
}

func (s *Server) handleReviewsGet(w http.ResponseWriter, r *http.Request) {
	// parse query parameters
	params := marshalQuery(r.URL.Query())
	id, ok := params["restaurant_id"]
	if !ok {
		console.Warning("on handleReviewsGet: need restaurant id")
		httpError(w, http.StatusBadRequest, pb.Code_PROTOCOL_MISMATCH)
		return
	}

	// translate restaurant id
	restID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		console.Warning("on handleReviewsGet: invalid restaurantd id; "+
			"id(%v): %v", id, err)
		httpError(w, http.StatusBadRequest, pb.Code_INVALID_DATA)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketReviewsGet{
		RestID: restID,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleReviewsGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleReviewsGet: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleReviewsGet: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	case server.ETNoSuchRestaurant:
		console.Warning(
			"on handleReviewsGet: restaurant not exists; RestID(%v)",
			dbReq.RestID)
		httpError(w, http.StatusBadRequest, pb.Code_RESTAURANT_NOT_EXISTS)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketReviews)
	if !ok {
		console.Warning("on handleReviewsGet: unexpected packet")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// copy reviews into response packet
	reviews := pb.Reviews{
		Reviews: make([]*pb.Review, 0, len(packet.Reviews)),
	}
	for _, r := range packet.Reviews {
		reviews.Reviews = append(reviews.Reviews,
			&pb.Review{
				ReviewId:  r.ID.Hex(),
				UserId:    r.UserID,
				UserName:  r.UserName,
				Score:     r.Score,
				Comment:   r.Comment,
				Menus:     r.Menus,
				Wait:      r.Wait,
				NumPeople: r.NumPeople,
				Timestamp: r.Timestamp,
			},
		)
	}

	// serialize user data
	ser, err := marshalResponse(r.Header, &reviews)
	if err != nil {
		console.Warning("on handleReviewsGet: failed to marshal review data")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// send review data for the restaurant
	baseHeader(w.Header())
	w.Write(ser)
	console.Info("send reviews; restaurant(%v) reviews(%v)",
		id, len(reviews.Reviews))
}

func (s *Server) handleRestaurant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleRestaurantGet(w, r)
	case "POST":
		s.handleRestaurantPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, pb.Code_METHOD_NOT_ALLOWED)
	}
}

func (s *Server) handleRestaurantGet(w http.ResponseWriter, r *http.Request) {
	// parse query parameters
	params := marshalQuery(r.URL.Query())
	id, ok := params["restaurant_id"]
	if !ok {
		console.Warning("on handleRestaurantGet: need restaurant id")
		httpError(w, http.StatusBadRequest, pb.Code_PROTOCOL_MISMATCH)
		return
	}

	// translate restaurant id
	restID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		console.Warning("on handleRestaurantGet: invalid restaurantd id; "+
			"id(%v): %v", id, err)
		httpError(w, http.StatusBadRequest, pb.Code_INVALID_DATA)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantGet{
		RestID: restID,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantGet: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantGet: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	case server.ETNoSuchRestaurant:
		console.Warning(
			"on handleRestaurantGet: restaurant not exists; RestID(%v)", restID)
		httpError(w, http.StatusBadRequest, pb.Code_RESTAURANT_NOT_EXISTS)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketRestaurant)
	if !ok {
		console.Warning("on handleRestaurantGet: unexpected packet")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// build response data
	rest := pb.Restaurant{
		Id:   packet.Restaurant.ID.Hex(),
		Name: packet.Restaurant.Name,
		Coord: &pb.Coordinate{
			Latitude:  packet.Restaurant.Coord.Latitude,
			Longitude: packet.Restaurant.Coord.Longitude,
		},
	}

	// serialize user data
	ser, err := marshalResponse(r.Header, &rest)
	if err != nil {
		console.Warning(
			"on handleRestaurantGet: failed to marshal restaurant data")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	baseHeader(w.Header())
	w.Write(ser)
	console.Info("sent restaurant data; restaurant(%v)", *packet)
}

func (s *Server) handleRestaurantPost(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq pb.RestaurantPost
	err := unmarshalBody(r.Header, r.Body, &userReq)
	if err != nil || userReq.Restaurant == nil {
		console.Warning("on handleRestaurantPost: failed to decode request")
		httpError(w, http.StatusBadRequest, pb.Code_PROTOCOL_MISMATCH)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantAdd{
		Name: userReq.Restaurant.Name,
		Coord: common.Coordinate{
			Latitude:  userReq.Restaurant.Coord.Latitude,
			Longitude: userReq.Restaurant.Coord.Longitude,
		},
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantPost: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantPost: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	console.Info("new restaurant added(%v)", userReq.Restaurant.Name)
}

func (s *Server) handleRestaurants(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleRestaurantsGet(w, r)
	case "POST":
		s.handleRestaurantsPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, pb.Code_METHOD_NOT_ALLOWED)
	}
}

func (s *Server) handleRestaurantsGet(w http.ResponseWriter, r *http.Request) {
	uid, _, err := s.authenticate(r.Header)
	if err != nil {
		console.Warning("on handleRestaurantsGet: %v", err)
		httpError(w, http.StatusBadRequest, pb.Code_AUTH_FAILED)
		return
	}

	// parse query parameters
	params := marshalQuery(r.URL.Query())
	var latitude, longitude float64
	if lat, ok := params["latitude"]; ok {
		latitude, err = strconv.ParseFloat(lat, 64)
		if err != nil {
			console.Warning("on handleRestaurantsGet: %v", err)
			httpError(w, http.StatusBadRequest, pb.Code_INVALID_DATA)
			return
		}
	}
	if lon, ok := params["longitude"]; ok {
		longitude, err = strconv.ParseFloat(lon, 64)
		if err != nil {
			console.Warning("on handleRestaurantsGet: %v", err)
			httpError(w, http.StatusBadRequest, pb.Code_INVALID_DATA)
			return
		}
	}
	if latitude == 0.0 || longitude == 0.0 {
		console.Warning("on handleRestaurantsGet: no latitude or longitude")
		httpError(w, http.StatusBadRequest, pb.Code_INVALID_DATA)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantsGet{
		UserID: uid,
		Coord: common.Coordinate{
			Latitude:  latitude,
			Longitude: longitude,
		},
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantsGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantsGet: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantsGet: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	case server.ETNoSuchUser:
		console.Warning(
			"on handleRestaurantsGet: user not exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest, pb.Code_USER_NOT_EXISTS)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketRestaurants)
	if !ok {
		console.Warning("on handleRestaurantsGet: unexpected packet")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// build response data
	rests := pb.Restaurants{
		Restaurants: make([]*pb.Restaurant, 0, len(packet.Restaurants)),
	}
	for _, r := range packet.Restaurants {
		rests.Restaurants = append(rests.Restaurants, &pb.Restaurant{
			Id:   r.ID.Hex(),
			Name: r.Name,
			Coord: &pb.Coordinate{
				Latitude:  r.Coord.Latitude,
				Longitude: r.Coord.Longitude,
			},
		})
	}

	// serialize user data
	ser, err := marshalResponse(r.Header, &rests)
	if err != nil {
		console.Warning(
			"on handleRestaurantsGet: failed to marshal restaurants data")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	baseHeader(w.Header())
	w.Write(ser)
	console.Info("sent restaurants data; count(%v)", len(rests.Restaurants))
}

func (s *Server) handleRestaurantsPost(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq pb.RestaurantsPost
	err := unmarshalBody(r.Header, r.Body, &userReq)
	if err != nil || userReq.Restaurants == nil {
		console.Warning("on handleRestaurantsPost: failed to decode request")
		httpError(w, http.StatusBadRequest, pb.Code_PROTOCOL_MISMATCH)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantsAdd{
		Restaurants: make([]*common.Restaurant, 0, len(userReq.Restaurants)),
	}
	for _, r := range userReq.Restaurants {
		if r.Coord == nil {
			console.Warning("on handleRestaurantsPost: invalid coordinate")
			httpError(w, http.StatusBadRequest, pb.Code_PROTOCOL_MISMATCH)
			return
		}

		dbReq.Restaurants = append(dbReq.Restaurants, &common.Restaurant{
			Name: r.Name,
			Coord: common.Coordinate{
				Latitude:  r.Coord.Latitude,
				Longitude: r.Coord.Longitude,
			},
		})
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantsPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantsPost: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantsPost: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	console.Info("new restaurants added; count(%v)", len(userReq.Restaurants))
}

func (s *Server) handleRanking(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleRankingGet(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed, pb.Code_METHOD_NOT_ALLOWED)
	}
}

func (s *Server) handleRankingGet(w http.ResponseWriter, r *http.Request) {
	// create packet for database server
	var dbReq = server.ADPacketRankingGet{}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRankingGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRankingGet: no packet received")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRankingGet: database internal error")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketUsers)
	if !ok {
		console.Warning("on handleRankingGet: unexpected packet")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	// copy user rankings for protobuf
	users := pb.Users{
		Users: make([]*pb.User, 0, len(packet.Users)),
	}
	for _, u := range packet.Users {
		users.Users = append(users.Users,
			&pb.User{
				Id:          u.UserID,
				Name:        u.Name,
				TotalExp:    u.Exp,
				ReviewCount: u.ReviewCount,
				LikeCount:   u.LikeCount,
			},
		)
	}

	// serialize user data
	ser, err := marshalResponse(r.Header, &users)
	if err != nil {
		console.Warning(
			"on handleRankingGet: failed to marshal ranking data")
		httpError(w, http.StatusInternalServerError, pb.Code_INTERNAL_ERROR)
		return
	}

	baseHeader(w.Header())
	w.Write(ser)
	console.Info("sent ranking data; count(%v)", len(users.Users))
}
