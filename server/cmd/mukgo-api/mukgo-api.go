package main

import (
	"flag"

	"github.com/isutare412/MukGo/server/api"
	"github.com/isutare412/MukGo/server/console"
)

func main() {
	// Handle flags
	addr := flag.String("addr", ":7777", "<ip:port> to run service")
	flag.Parse()

	// Create new api server
	server, err := api.NewServer("server", "server", "localhost:5672")
	if err != nil {
		console.Fatal("failed to create server: %v", err)
	}

	// Start service on port
	console.Info("start listen on %q...", *addr)
	if err := server.ListenAndServe(*addr); err != nil {
		console.Fatal("failed listen: %v", err)
	}
}
