package config_test

import "github.com/blazejsewera/notipie/producer/pkg/lib/config"

var exampleConfig = config.Config{
	EndpointConfig: config.EndpointConfig{
		Address: "localhost",
		Port:    1234,
	},
	AppID: "AppID",
}

const exampleConfigJSON = `{
    "endpointConfig": {
        "address": "localhost",
        "port": 1234
    },
    "appId": "AppID"
}
`

const exampleConfigYAML = `endpointConfig:
    address: localhost
    port: 1234
appId: AppID
`
