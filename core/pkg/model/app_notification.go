package model

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"gopkg.in/yaml.v3"
	"io"
)

type AppNotification struct {
	HashableNetNotification `yaml:",inline"`
	ID                      string `json:"id,omitempty" yaml:"id,omitempty"`
	ApiKey                  string `json:"apiKey,omitempty" yaml:"apiKey,omitempty"`
}

func (n AppNotification) ToJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return nil, fmt.Errorf("marshal AppNotification to JSON: %s", err)
	}
	return jsonBytes, nil
}

func (n AppNotification) ToYAML() ([]byte, error) {
	yamlBytes, err := yaml.Marshal(n)
	if err != nil {
		return nil, fmt.Errorf("marshal AppNotification to YAML: %s", err)
	}
	return yamlBytes, nil
}

func AddIDTo(n AppNotification) AppNotification {
	if n.ID != "" {
		return n
	}

	n.ID = hashAppNotification(n)
	return n
}

func hashAppNotification(n AppNotification) string {
	hashable := n.HashableNetNotification
	jsonBytes, err := json.Marshal(hashable)
	if err != nil {
		return ""
	}

	return uuid.GenerateBasedOnContent(jsonBytes)
}

func AppNotificationFromJSON(r io.Reader) (AppNotification, error) {
	d := json.NewDecoder(r)
	return appNotificationFromSerialized(d)
}

func AppNotificationFromYAML(r io.Reader) (AppNotification, error) {
	d := yaml.NewDecoder(r)
	return appNotificationFromSerialized(d)
}

func appNotificationFromSerialized(d Decoder) (AppNotification, error) {
	appNotification := AppNotification{}
	err := d.Decode(&appNotification)
	return appNotification, err
}

type Decoder interface {
	Decode(v any) error
}

func (n AppNotification) validate() bool {
	if n.AppName == "" || n.Title == "" {
		return false
	}
	return true
}

var ExampleAppNotification = AppNotification{
	HashableNetNotification: ExampleHashableNetNotification,
	ID:                      "frGOwBO6bNL/kbixYn3eJ6xS8WAewHK7qzt8q1cLVLs=",
	ApiKey:                  "ApiKey",
}

const ExampleAppNotificationJSON = `{
	"timestamp": "2022-06-14T22:22:22.000Z",
	"appName": "AppName",
	"appId": "AppID",
	"appImgUri": "AppImgURI",
	"title": "Title",
	"subtitle": "Subtitle",
	"body": "Body",
	"extUri": "ExtURI",
	"readUri": "ReadURI",
	"archiveUri": "ArchiveURI",
	"id": "frGOwBO6bNL/kbixYn3eJ6xS8WAewHK7qzt8q1cLVLs=",
	"apiKey": "ApiKey"
}`

const ExampleAppNotificationYAML = `timestamp: "2022-06-14T22:22:22.000Z"
appName: AppName
appId: AppID
appImgUri: AppImgURI
title: Title
subtitle: Subtitle
body: Body
extUri: ExtURI
readUri: ReadURI
archiveUri: ArchiveURI
id: frGOwBO6bNL/kbixYn3eJ6xS8WAewHK7qzt8q1cLVLs=
apiKey: ApiKey
`
