package main

import (
	"flag"

	"github.com/isutare412/MukGo/server/api"
	"github.com/isutare412/MukGo/server/log"
)

func main() {
	// Handle flags
	addr := flag.String("addr", ":7777", "<ip:port> to run service")
	flag.Parse()

	// Create new api server
	server := api.NewServer()

	// Start service on port
	log.Info("start listen on %q...", *addr)
	if err := server.ListenAndServe(*addr); err != nil {
		log.Fatal("failed listen: %v", err)
	}
}
