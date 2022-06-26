package config

import (
	"bytes"
	"encoding/json"
	"gopkg.in/yaml.v3"
)

func (c Config) ToJSON() ([]byte, error) {
	w := new(bytes.Buffer)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	err := enc.Encode(c)
	return w.Bytes(), err
}

func (c Config) ToYAML() ([]byte, error) {
	return yaml.Marshal(c)
}
