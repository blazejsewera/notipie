package model

import (
	"encoding/json"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
)

type HashableNetNotification struct {
	AppName    string `json:"appName" binding:"required"`
	AppImgUri  string `json:"appImgUri,omitempty"`
	Title      string `json:"title" binding:"required"`
	Subtitle   string `json:"subtitle,omitempty"`
	Body       string `json:"body,omitempty"`
	ExtUri     string `json:"extUri,omitempty"`
	ReadUri    string `json:"readUri,omitempty"`
	ArchiveUri string `json:"archiveUri,omitempty"`
}

type NetNotification struct {
	HashableNetNotification
	ID        string `json:"id,omitempty"`
	Timestamp string `json:"timestamp" binding:"required"`
	Read      bool   `json:"read,omitempty"`
}

func (n *NetNotification) AddID() {
	if n.ID != "" {
		return
	}

	hashable := n.HashableNetNotification
	jsonBytes, err := json.Marshal(hashable)
	if err != nil {
		return
	}

	n.ID = uuid.GenerateBasedOnContent(jsonBytes)
}
