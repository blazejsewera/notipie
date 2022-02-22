package model

import (
	"encoding/json"
	"fmt"
)

type ClientNotification struct {
	HashableNetNotification
	ID        string `json:"id"`
	Timestamp string `json:"timestamp"`
	Read      bool   `json:"read,omitempty"`
}

func (n ClientNotification) ToJSON() string {
	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func ClientNotificationFromJSON(jsonStr string) (ClientNotification, error) {
	clientNotification := ClientNotification{}
	err := json.Unmarshal([]byte(jsonStr), &clientNotification)
	if err != nil {
		return ClientNotification{}, err
	}
	if !clientNotification.validate() {
		return ClientNotification{}, fmt.Errorf(NotEnoughInfoInNotificationErrorMessage)
	}
	return clientNotification, nil
}

func (n ClientNotification) validate() bool {
	if n.AppName == "" || n.Title == "" || n.Timestamp == "" {
		return false
	}
	return true
}
