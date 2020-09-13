package main

import (
	"log"

	"github.com/isutare412/MukGo/server/api"
)

func main() {
	server := api.NewServer()

	// start listening on port
	log.Println("start listen...")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("on listen: %v", err)
	}
}
