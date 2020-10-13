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
	case "GET":
		s.handleUserGet(w, r)
	case "POST":
		s.handleUserPost(w, r)
	default:
		httpError(w, http.StatusMethodNotAllowed)
	}
}

func (s *Server) handleUserGet(w http.ResponseWriter, r *http.Request) {
	wait := make(chan struct{})

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

	response := func(success bool, p server.Packet) {
		defer func() {
			wait <- struct{}{}
		}()

		// failed to receive packet from database server
		if !success {
			console.Warning("on handleUserGet: no packet received")
			httpError(w, http.StatusInternalServerError)
			return
		}

		// check packet by type casting from interface
		var packet *server.DAPacketUser
		switch p.(type) {
		case *server.DAPacketUser:
			packet = p.(*server.DAPacketUser)
		case *server.DAPacketNoSuchUser:
			console.Warning("on handleUserGet: cannot find user; UserId(%v)",
				p.(*server.DAPacketNoSuchUser).UserID)
			httpError(w, http.StatusBadRequest)
			return
		default:
			console.Warning("on handleUserGet: unexpected packet")
			httpError(w, http.StatusInternalServerError)
			return
		}

		// calculate level
		level, residual, ratio := common.Exp2Level(packet.Exp)

		// serialize user data
		ser, err := json.Marshal(&ACUserInfo{
			Name:     packet.Name,
			Level:    level,
			TotalExp: packet.Exp,
			LevelExp: residual,
			ExpRatio: ratio,
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

	// send packet to database server and register response handler
	if err := s.send2DB(
		&dbReq,
		response,
	); err != nil {
		console.Warning("on handleUserGet: send2DB failed: %v", err)
		httpError(w, http.StatusInternalServerError)
		return
	}

	// wait for response
	<-wait
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
		switch p.(type) {
		case *server.DAPacketAck:
		case *server.DAPacketUserExist:
			console.Warning("on handleUserPost: user exists; UserId(%v)",
				p.(*server.DAPacketUserExist).UserID)
			httpError(w, http.StatusBadRequest)
			return
		default:
			console.Warning("on handleUserPost: unexpected packet")
			httpError(w, http.StatusInternalServerError)
			return
		}

		console.Info("new user created(%v)", userReq.Name)
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
		switch p.(type) {
		case *server.DAPacketAck:
		case *server.DAPacketError:
			console.Warning("on handleReviewPost: db error: %v",
				p.(*server.DAPacketError).Message)
			httpError(w, http.StatusInternalServerError)
			return
		default:
			console.Warning("on handleReviewPost: unexpected packet")
			httpError(w, http.StatusInternalServerError)
			return
		}

		console.Info("new review from user(%v)", userReq.UserID)
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
		switch p.(type) {
		case *server.DAPacketAck:
		case *server.DAPacketError:
			console.Warning("on handleRestaurantPost: db error: %v",
				p.(*server.DAPacketError).Message)
			httpError(w, http.StatusInternalServerError)
			return
		default:
			console.Warning("on handleRestaurantPost: unexpected packet")
			httpError(w, http.StatusInternalServerError)
			return
		}

		console.Info("new restaurant added(%v)", userReq.Name)
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
