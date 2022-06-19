package api

type Path struct {
	Path   string
	Schema string
}

const (
	HttpSchema = "http"
	WsSchema   = "ws"
)

var (
	Root          = Path{Path: "", Schema: HttpSchema}
	Push          = Path{Path: "push", Schema: HttpSchema}
	WebSocket     = Path{Path: "ws", Schema: WsSchema}
	Notifications = Path{Path: "notifications", Schema: HttpSchema}
)
