package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Nestik55/develop/dev11/api/server"
	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type Config struct {
	Server ServerConfig `yaml:"server"`
}

func initConfing() Config {
	data, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("error reading config file: %v", err)
	}

	var config Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarshalling yaml: %v", err)
	}

	return config
}

func main() {
	fmt.Println(time.Now())
	config := initConfing()

	server.Run(config.Server.Host, config.Server.Port)
}
