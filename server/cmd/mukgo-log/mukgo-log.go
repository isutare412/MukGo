package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/log"
	"github.com/joho/godotenv"
)

func main() {
	// load environment variables if env file is given
	if len(os.Args) >= 2 {
		if err := godotenv.Load(os.Args[1]); err != nil {
			console.Fatal("failed to load env file: %v", err)
		}
	}

	// build config from environment variables
	cfg, err := buildConfig()
	if err != nil {
		console.Fatal("failed to build config: %v", err)
	}

	// create new log server
	server, err := log.NewServer(cfg)
	if err != nil {
		console.Fatal("failed to create log server: %v", err)
	}

	// start log service
	console.Info("start log service...")
	if err := server.Run(); err != nil {
		console.Fatal("failed to run log server: %v", err)
	}
}

func buildConfig() (*log.ServerConfig, error) {
	rabbitPort, err := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		return nil, fmt.Errorf("on buildConfig: %v", err)
	}

	cfg := &log.ServerConfig{
		RabbitMQ: struct {
			User     string
			Password string
			IP       string
			Port     int
		}{
			os.Getenv("RABBITMQ_USER"),
			os.Getenv("RABBITMQ_PASSWORD"),
			os.Getenv("RABBITMQ_IP"),
			rabbitPort,
		},
	}

	return cfg, nil
}
