package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/isutare412/MukGo/server/api"
	"github.com/isutare412/MukGo/server/console"
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

	// create new api server
	server, err := api.NewServer(cfg)
	if err != nil {
		console.Fatal("failed to create api server: %v", err)
	}

	// start service on port
	addr := fmt.Sprintf("%s:%d", cfg.RestAPI.IP, cfg.RestAPI.Port)
	console.Info("start listen on %q...", addr)
	if err := server.ListenAndServe(addr); err != nil {
		console.Fatal("failed to run api server: %v", err)
	}
}

func buildConfig() (*api.ServerConfig, error) {
	rabbitPort, err := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		return nil, fmt.Errorf("on buildConfig: %v", err)
	}

	servPort, err := strconv.Atoi(os.Getenv("SERVICE_PORT"))
	if err != nil {
		return nil, fmt.Errorf("on buildConfig: %v", err)
	}

	cfg := &api.ServerConfig{
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
		RestAPI: struct {
			IP   string
			Port int
		}{
			os.Getenv("SERVICE_IP"),
			servPort,
		},
	}

	return cfg, nil
}
