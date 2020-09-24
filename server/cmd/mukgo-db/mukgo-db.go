package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/db"
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

	// create new database server
	server, err := db.NewServer(cfg)
	if err != nil {
		console.Fatal("failed to create database server: %v", err)
	}

	// initialize database
	if err := server.InitDB(); err != nil {
		console.Fatal("failed to initiate database server: %v", err)
	}

	// start database service
	console.Info("start database service...")
	if err := server.Run(); err != nil {
		console.Fatal("failed to run database server: %v", err)
	}
}

func buildConfig() (*db.ServerConfig, error) {
	rabbitPort, err := strconv.Atoi(os.Getenv("RABBITMQ_PORT"))
	if err != nil {
		return nil, fmt.Errorf("on buildConfig: %v", err)
	}

	mongoPort, err := strconv.Atoi(os.Getenv("MONGODB_PORT"))
	if err != nil {
		return nil, fmt.Errorf("on buildConfig: %v", err)
	}

	cfg := &db.ServerConfig{
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
		MongoDB: struct {
			User     string
			Password string
			IP       string
			Port     int
		}{
			os.Getenv("MONGODB_USER"),
			os.Getenv("MONGODB_PASSWORD"),
			os.Getenv("MONGODB_IP"),
			mongoPort,
		},
	}

	return cfg, nil
}
