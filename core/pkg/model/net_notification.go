package model

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

const NotEnoughInfoInNotificationErrorMessage = "not enough information in net notification"

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
