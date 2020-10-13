package api

import (
	"encoding/json"
	"net/http"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/common"
	"github.com/isutare412/MukGo/server/console"
)

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		s.handleUserPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleUserPost(w http.ResponseWriter, r *http.Request) {
	wait := make(chan struct{})

	// parse request from client
	var userReq CAUserPost
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handlerUserPost: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketUserAdd{
		UserID: userReq.UserID,
		Name:   userReq.Name,
	}

	response := func(success bool, p server.Packet) {
		defer func() {
			wait <- struct{}{}
		}()

		// failed to receive packet from database server
		if !success {
			console.Warning("on handlerUserPost: no packet received")
			httpError(w, http.StatusInternalServerError)
			return
		}

		// check packet by type casting from interface
		_, ok := p.(*server.DAPacketAck)
		if !ok {
			console.Warning("on handlerUserPost: failed to write to database")
			httpError(w, http.StatusConflict)
			return
		}

		console.Info("new user created(%s)", userReq.Name)
	}

	// send packet to database server and register response handler
	if err := s.send2DB(
		&dbReq,
		response,
	); err != nil {
		console.Warning("on handlerUserPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response
	<-wait
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
	wait := make(chan struct{})

	// parse request from client
	var userReq CAReviewPost
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		console.Warning("on handlerReviewPost: failed to decode request")
		httpError(w, http.StatusBadRequest)
		return
	}

	// create packet for database server
	var dbReq = server.ADPacketReviewAdd{
		UserID:  userReq.UserID,
		Score:   userReq.Score,
		Comment: userReq.Comment,
	}

	response := func(success bool, p server.Packet) {
		defer func() {
			wait <- struct{}{}
		}()

		// failed to receive packet from database server
		if !success {
			console.Warning("on handlerReviewPost: no packet received")
			httpError(w, http.StatusInternalServerError)
			return
		}

		// check packet by type casting from interface
		_, ok := p.(*server.DAPacketAck)
		if !ok {
			console.Warning("on handlerReviewPost: failed to write to database")
			httpError(w, http.StatusInternalServerError)
			return
		}

		console.Info("new review from user(%d)", userReq.UserID)
	}

	// send packet to database server and register response handler
	if err := s.send2DB(
		&dbReq,
		response,
	); err != nil {
		console.Warning("on handlerReviewPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response
	<-wait
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
	wait := make(chan struct{})

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

	response := func(success bool, p server.Packet) {
		defer func() {
			wait <- struct{}{}
		}()

		// failed to receive packet from database server
		if !success {
			console.Warning("on handleRestaurantPost: no packet received")
			httpError(w, http.StatusInternalServerError)
			return
		}

		// check packet by type casting from interface
		_, ok := p.(*server.DAPacketAck)
		if !ok {
			console.Warning(
				"on handleRestaurantPost: failed to write to database")
			httpError(w, http.StatusInternalServerError)
			return
		}

		console.Info("new restaurant added(%s)", userReq.Name)
	}

	// send packet to database server and register response handler
	if err := s.send2DB(
		&dbReq,
		response,
	); err != nil {
		console.Warning("on handleRestaurantPost: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response
	<-wait
}
