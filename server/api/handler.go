package api

import (
	"encoding/json"
	"net/http"

	"github.com/isutare412/MukGo/server"
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
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleUserGet(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq CAUserGet
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handleUserGet: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketUserGet{
		UserID: userReq.UserID,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleUserGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleUserGet: no packet received")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleUserGet: database internal error")
		httpError(w, http.StatusInternalServerError)
		return
	case server.ETNoSuchUser:
		console.Warning("on handleUserGet: user not exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketUser)
	if !ok {
		console.Warning("on handleUserGet: unexpected packet")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// calculate level
	level, levExp, curExp, ratio := common.Exp2Level(packet.Exp)
	sightRadius := common.Level2Sight(level)

	// serialize user data
	ser, err := json.Marshal(&ACUserInfo{
		Name:        packet.Name,
		Level:       level,
		TotalExp:    packet.Exp,
		LevelExp:    levExp,
		CurExp:      curExp,
		ExpRatio:    ratio,
		SightRadius: sightRadius,
	})
	if err != nil {
		console.Warning("on handleUserGet: failed to marshal user data")
		httpError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ser)
	console.Info("sent user data; User(%v)", *packet)
}

func (s *Server) handleUserPost(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq CAUserPost
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handleUserPost: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketUserAdd{
		UserID: userReq.UserID,
		Name:   userReq.Name,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleUserPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleUserPost: no packet received")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleUserPost: database internal error")
		httpError(w, http.StatusInternalServerError)
		return
	case server.ETUserExists:
		console.Warning(
			"on handleUserPost: user already exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest)
		return
	}

	console.Info("new user created(%v)", userReq.Name)
}

func (s *Server) handleReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.handleReviewPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleReviewPost(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq CAReviewPost
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handleReviewPost: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	restID, err := primitive.ObjectIDFromHex(userReq.RestID)
	if err != nil {
		console.Warning("on handleReviewPost: invalid restaurantd id; "+
			"id(%v): %v", userReq.RestID, err)
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketReviewAdd{
		UserID:  userReq.UserID,
		RestID:  restID,
		Score:   userReq.Score,
		Comment: userReq.Comment,
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleReviewPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleReviewPost: no packet received")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleReviewPost: database internal error")
		httpError(w, http.StatusInternalServerError)
		return
	case server.ETNoSuchUser:
		console.Warning("on handleReviewPost: user not exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest)
		return
	case server.ETNoSuchRestaurant:
		console.Warning(
			"on handleReviewPost: restaurant not exists; RestID(%v)",
			dbReq.RestID)
		httpError(w, http.StatusBadRequest)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketUser)
	if !ok {
		console.Warning("on handleReviewPost: unexpected packet")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// calculate level
	level, levExp, curExp, ratio := common.Exp2Level(packet.Exp)
	sightRadius := common.Level2Sight(level)

	// serialize user data
	ser, err := json.Marshal(&ACUserInfo{
		Name:        packet.Name,
		Level:       level,
		TotalExp:    packet.Exp,
		LevelExp:    levExp,
		CurExp:      curExp,
		ExpRatio:    ratio,
		SightRadius: sightRadius,
	})
	if err != nil {
		console.Warning("on handleReviewPost: failed to marshal user data")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// send updated user data
	w.Header().Set("Content-Type", "application/json")
	w.Write(ser)
	console.Info("new review from user; User(%v)", *packet)
}

func (s *Server) handleRestaurant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.handleRestaurantPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleRestaurantPost(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq CARestaurantPost
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handleRestaurantPost: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantAdd{
		Name: userReq.Name,
		Coord: common.Coordinate{
			Latitude:  userReq.Latitude,
			Longitude: userReq.Longitude,
		},
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantPost: no packet received")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantPost: database internal error")
		httpError(w, http.StatusInternalServerError)
		return
	}

	console.Info("new restaurant added(%v)", userReq.Name)
}

func (s *Server) handleRestaurants(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.handleRestaurantsGet(w, r)
	case "POST":
		s.handleRestaurantsPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleRestaurantsGet(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq CARestaurantsGet
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handleRestaurantsGet: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantsGet{
		UserID: userReq.UserID,
		Coord: common.Coordinate{
			Latitude:  userReq.Latitude,
			Longitude: userReq.Longitude,
		},
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantsGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantsGet: no packet received")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantsGet: database internal error")
		httpError(w, http.StatusInternalServerError)
		return
	case server.ETNoSuchUser:
		console.Warning(
			"on handleRestaurantsGet: user not exists; UserID(%v)",
			dbReq.UserID)
		httpError(w, http.StatusBadRequest)
		return
	}

	// check packet by type casting from interface
	packet, ok := p.(*server.DAPacketRestaurants)
	if !ok {
		console.Warning("on handleRestaurantsGet: unexpected packet")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// build response data
	rests := ACRestaurantsInfo{
		Restaurants: make([]*Restaurant, 0, len(packet.Restaurants)),
	}
	for _, r := range packet.Restaurants {
		rests.Restaurants = append(rests.Restaurants, &Restaurant{
			ID:        r.ID.Hex(),
			Name:      r.Name,
			Latitude:  r.Coord.Latitude,
			Longitude: r.Coord.Longitude,
		})
	}

	// serialize user data
	ser, err := json.Marshal(&rests)
	if err != nil {
		console.Warning(
			"on handleRestaurantsGet: failed to marshal restaurants data")
		httpError(w, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(ser)
	console.Info("sent restaurants data; count(%v)", len(rests.Restaurants))
}

func (s *Server) handleRestaurantsPost(w http.ResponseWriter, r *http.Request) {
	// parse request from client
	var userReq CARestaurantsPost
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handleRestaurantsPost: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketRestaurantsAdd{
		Restaurants: make([]*common.Restaurant, 0, len(userReq.Restaurants)),
	}
	for _, r := range userReq.Restaurants {
		dbReq.Restaurants = append(dbReq.Restaurants, &common.Restaurant{
			Name: r.Name,
			Coord: common.Coordinate{
				Latitude:  r.Latitude,
				Longitude: r.Longitude,
			},
		})
	}

	// send packet to database server and register response handler
	response, err := s.send2DB(&dbReq)
	if err != nil {
		console.Warning("on handleRestaurantsPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response packet
	p := <-response

	// failed to receive packet from database server
	if p == nil {
		console.Warning("on handleRestaurantsPost: no packet received")
		httpError(w, http.StatusInternalServerError)
		return
	}

	// handle error packet
	switch getError(p) {
	case server.ETInvalid:
		break // not error
	case server.ETInternal:
		console.Warning("on handleRestaurantsPost: database internal error")
		httpError(w, http.StatusInternalServerError)
		return
	}

	console.Info("new restaurants added; count(%v)",
		len(userReq.Restaurants))
}
