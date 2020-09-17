package main

import (
	"flag"

	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/log"
)

func main() {
	// handle flags
	mqaddr := flag.String(
		"mqaddr", "localhost:5672", "<ip:port> of RabbitMQ to connect")
	flag.Parse()

	// create new api server
	server, err := log.NewServer("server", "server", *mqaddr)
	if err != nil {
		console.Fatal("failed to create log server: %v", err)
	}

	// start log service
	console.Info("start service...")
	if err := server.Run(); err != nil {
		console.Fatal("failed to run log server: %v", err)
	}
}
