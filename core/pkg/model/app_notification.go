package model

import (
	"encoding/json"
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"io"
)

type AppNotification struct {
	HashableNetNotification
	ID        string `json:"id,omitempty"`
	Timestamp string `json:"timestamp"`
	ApiKey    string `json:"apiKey,omitempty"`
}

func (n AppNotification) ToJSON() ([]byte, error) {
	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return nil, fmt.Errorf("marshal AppNotification: %s", err)
	}
	return jsonBytes, nil
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

func AppNotificationFromJSON(jsonBytes []byte) (AppNotification, error) {
	appNotification := AppNotification{}
	err := json.Unmarshal(jsonBytes, &appNotification)
	if err != nil {
		return AppNotification{}, err
	}
	if !appNotification.validate() {
		return AppNotification{}, fmt.Errorf(NotEnoughInfoInNotificationErrorMessage)
	}
	return appNotification, nil
}

func AppNotificationFromReader(r io.Reader) (AppNotification, error) {
	appNotification := AppNotification{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(&appNotification)
	if err != nil {
		return AppNotification{}, err
	}
	if !appNotification.validate() {
		return AppNotification{}, fmt.Errorf(NotEnoughInfoInNotificationErrorMessage)
	}
	return appNotification, nil
}

func (n AppNotification) validate() bool {
	if n.AppName == "" || n.Title == "" || n.Timestamp == "" {
		return false
	}
	return true
}

var ExampleAppNotification = AppNotification{
	HashableNetNotification: ExampleHashableNetNotification,
	ID:                      "frGOwBO6bNL/kbixYn3eJ6xS8WAewHK7qzt8q1cLVLs=",
	Timestamp:               "2022-06-14T22:22:22.000Z",
	ApiKey:                  "ApiKey",
}

const ExampleAppNotificationJSON = `{
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
	"timestamp": "2022-06-14T22:22:22.000Z",
	"apiKey": "ApiKey"
}`
