package main

import (
	"flag"
	"log"

	"github.com/isutare412/MukGo/server/api"
)

func main() {
	// Handle flags
	addr := flag.String("addr", ":7777", "<ip:port> to run service")
	flag.Parse()

	// Create new api server
	server := api.NewServer()

	// Start service on port
	log.Printf("start listen on %q...", *addr)
	if err := server.ListenAndServe(*addr); err != nil {
		log.Fatalf("on listen: %v", err)
	}
}
