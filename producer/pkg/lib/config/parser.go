package config

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
)

func FromJSON(r io.Reader) (Config, error) {
	d := json.NewDecoder(r)
	return fromSerialized(d)
}

func FromYAML(r io.Reader) (Config, error) {
	d := yaml.NewDecoder(r)
	return fromSerialized(d)
}

func fromSerialized(d Decoder) (Config, error) {
	config := Config{}
	err := d.Decode(&config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

type Decoder interface {
	Decode(v any) error
}
