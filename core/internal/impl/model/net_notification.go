package model

import (
	"encoding/json"
	"fmt"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
)

type HashableNetNotification struct {
	AppName    string `json:"appName"`
	AppID      string `json:"appId,omitempty"`
	AppImgURI  string `json:"appImgUri,omitempty"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle,omitempty"`
	Body       string `json:"body,omitempty"`
	ExtURI     string `json:"extUri,omitempty"`
	ReadURI    string `json:"readUri,omitempty"`
	ArchiveURI string `json:"archiveUri,omitempty"`
}

type NetNotification struct {
	HashableNetNotification
	ID        string `json:"id,omitempty"`
	Timestamp string `json:"timestamp"`
	Read      bool   `json:"read,omitempty"`
	ApiKey    string `json:"apiKey,omitempty"`
}

func (n NetNotification) ToJSON() string {
	jsonBytes, err := json.Marshal(n)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func AddIDTo(n NetNotification) NetNotification {
	if n.ID != "" {
		return n
	}

	n.ID = hashNetNotification(n)
	return n
}

func hashNetNotification(n NetNotification) string {
	hashable := n.HashableNetNotification
	jsonBytes, err := json.Marshal(hashable)
	if err != nil {
		return ""
	}

	return uuid.GenerateBasedOnContent(jsonBytes)
}

func NetNotificationFromJSON(jsonStr string) (NetNotification, error) {
	netNotification := NetNotification{}
	err := json.Unmarshal([]byte(jsonStr), &netNotification)
	if err != nil {
		return NetNotification{}, err
	}
	if !validateNetNotificationFields(netNotification) {
		return NetNotification{}, fmt.Errorf(NotEnoughInfoInNotificationErrorMessage)
	}
	return netNotification, nil
}

func validateNetNotificationFields(n NetNotification) bool {
	if n.AppName == "" || n.Title == "" || n.Timestamp == "" {
		return false
	}
	return true
}

const NotEnoughInfoInNotificationErrorMessage = "not enough information in net notification"
