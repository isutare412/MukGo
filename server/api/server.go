package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/mq"
)

// NewServer creates Server struct safely.
func NewServer(mqid, mqpw, mqaddr string) (*Server, error) {
	var server = &Server{
		mux: http.NewServeMux(),
	}

	// register api handlers
	server.registerHandlers()
	console.Info("registered handlers")

	// establish rabbitmq session
	session, err := mq.NewSession(mqid, mqpw, mqaddr)
	if err != nil {
		return nil, fmt.Errorf("on NewServer: %v", err)
	}
	server.mqss = session
	console.Info("session(%q) established between RabbitMQ", mqaddr)

	// initialize message queues
	err = server.initQueue()
	if err != nil {
		return nil, fmt.Errorf("on Newserver: %v", err)
	}

	return server, nil
}

// Server runs as API server for MukGo service. Server should be created with
// NewServer function.
type Server struct {
	mux  *http.ServeMux
	mqss *mq.Session
}

// ListenAndServe starts Server. If addr is blank, ":http" is used, which
// uses 80 port.
func (s *Server) ListenAndServe(addr string) error {
	if addr == "" {
		addr = ":http"
	}

	// Assign ServeMux to Server.
	httpserver := &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	return httpserver.ListenAndServe()
}

func (s *Server) registerHandlers() {
	s.mux.HandleFunc("/devel", s.handlerDevel)
}

func (s *Server) handlerDevel(w http.ResponseWriter, r *http.Request) {
	// Parse request into RestRequest.
	var req RestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		console.Warning("failed to decode request: %v", err)
		return
	}

	console.Info("message from %q: %q", req.User, req.Message)

	// marshal response into byte slice
	res := RestResponse{"Hello, Client!"}
	resBytes, err := json.Marshal(res)
	if err != nil {
		console.Warning("failed to encode response: %v", err)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}
