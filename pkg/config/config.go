package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	App
	Environment string
}

type App struct {
	Server  server  `json:"server"`
	Mongodb mongodb `json:"mongodb"`
}

type server struct {
	Port string `json:"port"`
}

type mongodb struct {
	DatabaseName               string `json:"database_name"`
	CollectionName             string `json:"collection_name"`
	Url                        string `json:"url"`
	DisconnectTimeoutInSeconds int    `json:"disconnect_timeout_in_s"`
}

// Initializes a Configuration
func NewConfiguration() Configuration {
	env := os.Getenv("ENVIRONMENT")
	if env == "" {
		env = "local"
	}
	return Configuration{
		Environment: env,
		App:         readApplicationConfig(env),
	}
}

// Reads configuration from a json file
func readApplicationConfig(env string) App {
	var app App

	f, err := os.Open("./pkg/config/support/" + env + "/config.json")

	defer f.Close()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	jsonParser := json.NewDecoder(f)

	if err = jsonParser.Decode(&app); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return app
}
