package impl

type Notification struct {
	AppName    string `json:"appName"`
	Timestamp  string `json:"timestamp"`
	AppImgUri  string `json:"appImgUri,omitempty"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle,omitempty"`
	Body       string `json:"body,omitempty"`
	ID         string `json:"id,omitempty"`
	Read       bool   `json:"read,omitempty"`
	ExtUri     string `json:"extUri,omitempty"`
	ReadUri    string `json:"readUri,omitempty"`
	ArchiveUri string `json:"archiveUri,omitempty"`
}
