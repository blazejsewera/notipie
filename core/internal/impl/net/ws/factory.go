package ws

import "github.com/gorilla/websocket"

type HubFactory interface {
	GetHub() Hub
}

type HubFactoryFunc func() Hub

func (f HubFactoryFunc) GetHub() Hub {
	return f()
}

var DefaultHubFactory = HubFactoryFunc(func() Hub {
	return NewHub(NewClientFactory())
})

type ClientFactory interface {
	SetHub(hub Hub)
	GetClient(conn *websocket.Conn) Client
}

type DefaultClientFactory struct {
	hub Hub
}

func NewClientFactory() *DefaultClientFactory {
	return &DefaultClientFactory{}
}

// DefaultClientFactory implements interfaces below
var _ ClientFactory = (*DefaultClientFactory)(nil)

func (f *DefaultClientFactory) SetHub(hub Hub) {
	f.hub = hub
}

func (f *DefaultClientFactory) GetClient(conn *websocket.Conn) Client {
	return NewClient(f.hub, conn)
}
