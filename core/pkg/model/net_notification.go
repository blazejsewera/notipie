package model

type HashableNetNotification struct {
	AppName    string `json:"appName" yaml:"appName"`
	AppID      string `json:"appId,omitempty" yaml:"appId,omitempty"`
	AppImgURI  string `json:"appImgUri,omitempty" yaml:"appImgUri,omitempty"`
	Title      string `json:"title" yaml:"title"`
	Subtitle   string `json:"subtitle,omitempty" yaml:"subtitle,omitempty"`
	Body       string `json:"body,omitempty" yaml:"body,omitempty"`
	ExtURI     string `json:"extUri,omitempty" yaml:"extUri,omitempty"`
	ReadURI    string `json:"readUri,omitempty" yaml:"readUri,omitempty"`
	ArchiveURI string `json:"archiveUri,omitempty" yaml:"archiveUri,omitempty"`
}

const NotEnoughInfoInNotificationErrorMessage = "not enough information in net notification: missing appName, title, and/or timestamp"

var ExampleHashableNetNotification = HashableNetNotification{
	AppName:    "AppName",
	AppID:      "AppID",
	AppImgURI:  "AppImgURI",
	Title:      "Title",
	Subtitle:   "Subtitle",
	Body:       "Body",
	ExtURI:     "ExtURI",
	ReadURI:    "ReadURI",
	ArchiveURI: "ArchiveURI",
}
