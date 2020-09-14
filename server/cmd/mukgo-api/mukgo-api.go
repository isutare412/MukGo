package main

import (
	"log"

	"github.com/isutare412/MukGo/server/api"
)

func main() {
	server := api.NewServer()

	// start service on port
	log.Println("start listen...")
	if err := server.ListenAndServe(":7777"); err != nil {
		log.Fatalf("on listen: %v", err)
	}
}
