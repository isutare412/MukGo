package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// NewServer creates Server struct safely.
func NewServer() *Server {
	var server = &Server{
		hs: &http.Server{
			Addr: ":7777",
		},
		mux: http.NewServeMux(),
	}

	// Register handlers.
	server.mux.HandleFunc("/devel", server.handlerDevel)

	// Assign ServeMux to Server.
	server.hs.Handler = server.mux

	return server
}

// Server runs as API server for MukGo service. Server should be created with
// NewServer function.
type Server struct {
	hs  *http.Server
	mux *http.ServeMux
}

// ListenAndServe starts Server.
func (s *Server) ListenAndServe() error {
	return s.hs.ListenAndServe()
}

func (s *Server) handlerDevel(w http.ResponseWriter, r *http.Request) {
	// Parse request into RestRequest.
	var req RestRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("on decode: %v", err)
		return
	}

	log.Printf("message from %q: %q", req.User, req.Message)

	// marshal response into byte slice
	res := RestResponse{"Hello, Client!"}
	resBytes, err := json.Marshal(res)
	if err != nil {
		log.Printf("on marshal response: %v", err)
		return
	}

	// send response
	w.Header().Set("Content-Type", "application/json")
	w.Write(resBytes)
}
