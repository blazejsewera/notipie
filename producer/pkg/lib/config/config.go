package config

type Config struct {
	EndpointConfig `json:"endpointConfig" yaml:"endpointConfig"`
	AppID          string `json:"appId,omitempty" yaml:"appId,omitempty"`
}

type EndpointConfig struct {
	Address string `json:"address" yaml:"address"`
	Port    int    `json:"port" yaml:"port"`
}
