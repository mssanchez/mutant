package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration holds all the configurations needed to run the application
type Configuration struct {
	App
	Environment string
}

type App struct {
	Server  Server  `json:"server"`
	Mongodb Mongodb `json:"mongodb"`
}

type Server struct {
	Port string `json:"port"`
}

// Mongodb holds configuration needed to connect to to MongoDB
type Mongodb struct {
	DatabaseName               string `json:"database_name"`
	CollectionName             string `json:"collection_name"`
	Url                        string `json:"url"`
	Password                   string
	DisconnectTimeoutInSeconds int `json:"disconnect_timeout_in_s"`
}

// NewConfiguration initializes a Configuration
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

// readApplicationConfig reads configuration from a json file
func readApplicationConfig(env string) App {
	var app App
	var f *os.File
	var err error

	switch {
	case env == "test":
		f, err = os.Open("./support/" + env + "/config.json")
	default:
		f, err = os.Open("./pkg/config/support/" + env + "/config.json")
	}

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer f.Close()

	jsonParser := json.NewDecoder(f)

	if err = jsonParser.Decode(&app); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return app
}
