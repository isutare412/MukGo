package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/isutare412/MukGo/server/api"
	"github.com/isutare412/MukGo/server/console"
	"gopkg.in/yaml.v2"
)

func main() {
	if len(os.Args) < 2 {
		console.Fatal("need yaml file for configuration")
	}

	// read settings from yaml file
	fileName := os.Args[1]
	fileBody, err := ioutil.ReadFile(fileName)
	if err != nil {
		console.Fatal("cannot open file: %q", fileName)
	}

	// parse yaml file
	var cfg api.ServerConfig
	if err := yaml.Unmarshal(fileBody, &cfg); err != nil {
		console.Fatal("failed to parse file: %q", fileName)
	}

	// create new api server
	server, err := api.NewServer(&cfg)
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
