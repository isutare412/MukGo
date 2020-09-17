package main

import (
	"flag"

	"github.com/isutare412/MukGo/server/api"
	"github.com/isutare412/MukGo/server/console"
)

func main() {
	// handle flags
	addr := flag.String("addr", ":7777", "<ip:port> to run service")
	mqaddr := flag.String(
		"mqaddr", "localhost:5672", "<ip:port> of RabbitMQ to connect")
	flag.Parse()

	// create new api server
	server, err := api.NewServer("server", "server", *mqaddr)
	if err != nil {
		console.Fatal("failed to create api server: %v", err)
	}

	// start service on port
	console.Info("start listen on %q...", *addr)
	if err := server.ListenAndServe(*addr); err != nil {
		console.Fatal("failed to run api server: %v", err)
	}
}
