package api

import (
	"encoding/json"
	"net/http"

	"github.com/isutare412/MukGo/server"
	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/loc"
)

func (s *Server) handleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		wait := make(chan struct{})

		// parse request from client
		var userReq JSONUserPost
		if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
			console.Warning("on handlerUser: failed to decode request")
			httpError(w, http.StatusBadRequest)
			return
		}

		// create packet for database server
		var dbReq = server.PacketUserAdd{
			UserID: userReq.UserID,
			Name:   userReq.Name,
		}

		response := func(success bool, p server.Packet) {
			defer func() {
				wait <- struct{}{}
			}()

			// failed to receive packet from database server
			if !success {
				console.Warning("on handlerUser: no packet received")
				httpError(w, http.StatusInternalServerError)
				return
			}

			// check packet by type casting from interface
			_, ok := p.(*server.PacketAck)
			if !ok {
				console.Warning("on handlerUser: failed to write to database")
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
			console.Warning("send2DB failed: %v", err)
			httpError(w, http.StatusInternalServerError)
			return
		}

		// wait for response
		<-wait

	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleReview(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		wait := make(chan struct{})

		// parse request from client
		var userReq JSONReviewPost
		if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
			console.Warning("on handlerReview: failed to decode request")
			httpError(w, http.StatusBadRequest)
			return
		}

		// create packet for database server
		var dbReq = server.PacketReviewAdd{
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
				console.Warning("on handlerReview: no packet received")
				httpError(w, http.StatusInternalServerError)
				return
			}

			// check packet by type casting from interface
			_, ok := p.(*server.PacketAck)
			if !ok {
				console.Warning("on handlerReview: failed to write to database")
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
			console.Warning("send2DB failed: %v", err)
			httpError(w, http.StatusInternalServerError)
			return
		}

		// wait for response
		<-wait

	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleRestaurant(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		wait := make(chan struct{})

		// parse request from client
		var userReq JSONRestaurantPost
		if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
			console.Warning("on handlerRestaurant: failed to decode request")
			httpError(w, http.StatusBadRequest)
			return
		}

		// create packet for database server
		var dbReq = server.PacketRestaurantAdd{
			Name: userReq.Name,
			Coord: loc.Coordinate{
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
				console.Warning("on handleRestaurant: no packet received")
				httpError(w, http.StatusInternalServerError)
				return
			}

			// check packet by type casting from interface
			_, ok := p.(*server.PacketAck)
			if !ok {
				console.Warning("on handleRestaurant: failed to write to database")
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
			console.Warning("send2DB failed: %v", err)
			httpError(w, http.StatusInternalServerError)
			return
		}

		// wait for response
		<-wait

	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}
