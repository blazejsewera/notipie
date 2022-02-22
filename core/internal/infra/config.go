package infra

import (
	"encoding/json"
	"fmt"
)

type Config struct {
	prod    bool
	address string
	port    int
}

func DefaultConfig() Config {
	return Config{
		prod:    false,
		address: "localhost",
		port:    8080,
	}
}

func ConfigFromJSON(jsonStr string) Config {
	config := Config{}
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		panic(fmt.Sprintf("cannot unmarshal config json: %s", jsonStr))
	}
	return config
}
