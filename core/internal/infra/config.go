package infra

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Config struct {
	Prod    bool   `json:"prod"`
	Address string `json:"address"`
	Port    int    `json:"port"`
}

func DefaultConfig() Config {
	configJSON := readDefaultConfigFile()
	return ConfigFromJSON(configJSON)
}

func readDefaultConfigFile() string {
	bytes, err := ioutil.ReadFile(DefaultConfigFilename)
	if err != nil {
		panic(fmt.Sprint("could not read default config: ", DefaultConfigFilename))
	}
	return string(bytes)
}

func ConfigFromJSON(jsonStr string) Config {
	config := Config{}
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		panic(fmt.Sprintf("cannot unmarshal config json: %s", jsonStr))
	}
	return config
}

const DefaultConfigFilename = "notipie.config.json"
