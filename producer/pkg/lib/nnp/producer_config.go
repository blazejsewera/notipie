package nnp

import "net/url"

type ProducerConfig struct {
	AppID    string
	Endpoint ProducerEndpointConfig
}

type ProducerEndpointConfig struct {
	RootURL url.URL
	PushURL url.URL
}

type AppIDSaver interface {
	SaveAppID(appID string) error
}
