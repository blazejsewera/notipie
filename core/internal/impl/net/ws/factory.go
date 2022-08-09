package ws

import "github.com/gorilla/websocket"

type HubFactory interface {
	GetHub() Hub
}

type HubFactoryFunc func() Hub

func (f HubFactoryFunc) GetHub() Hub {
	return f()
}

var WebSocketHubFactory = HubFactoryFunc(func() Hub {
	return NewHubImpl(NewClientFactory())
})

type ClientFactory interface {
	SetHub(hub Hub)
	GetClient(conn *websocket.Conn) Client
}

type WebSocketClientFactory struct {
	hub Hub
}

func NewClientFactory() *WebSocketClientFactory {
	return &WebSocketClientFactory{}
}

// @impl
var _ ClientFactory = (*WebSocketClientFactory)(nil)

func (f *WebSocketClientFactory) SetHub(hub Hub) {
	f.hub = hub
}

func (f *WebSocketClientFactory) GetClient(conn *websocket.Conn) Client {
	return NewClient(f.hub, conn)
}
