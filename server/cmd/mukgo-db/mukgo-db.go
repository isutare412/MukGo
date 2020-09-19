package main

import (
	"io/ioutil"
	"os"

	"github.com/isutare412/MukGo/server/console"
	"github.com/isutare412/MukGo/server/db"
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
	var cfg db.ServerConfig
	if err := yaml.Unmarshal(fileBody, &cfg); err != nil {
		console.Fatal("failed to parse file: %q", fileName)
	}

	// create new database server
	_, err = db.NewServer(&cfg)
	if err != nil {
		console.Fatal("failed to create log server: %v", err)
	}

	// start database service
	console.Info("start service...")
}
