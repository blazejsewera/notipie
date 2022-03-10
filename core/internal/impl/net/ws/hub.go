package ws

import (
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ClientHub interface {
	Broadcast(notification model.ClientNotification)
	Register(conn *websocket.Conn)
	Unregister(clientUUID string)
}

type ClientHubFactory interface {
	GetClientHub() ClientHub
}

type ClientHubFactoryFunc func() ClientHub

func (f ClientHubFactoryFunc) GetClientHub() ClientHub {
	return f()
}

var DefaultClientHubFactory = ClientHubFactoryFunc(func() ClientHub {
	return NewHub()
})

type Hub struct {
	clients    map[string]*Client
	broadcast  chan model.ClientNotification
	register   chan *websocket.Conn
	unregister chan string
	l          *zap.Logger
}

// Hub implements interfaces below
var _ ClientHub = (*Hub)(nil)
var _ util.Starter = (*Hub)(nil)

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan model.ClientNotification),
		register:   make(chan *websocket.Conn),
		unregister: make(chan string),
		clients:    make(map[string]*Client),
		l:          log.For("impl").Named("net").Named("hub"),
	}
}

func (h *Hub) Broadcast(notification model.ClientNotification) {
	h.broadcast <- notification
}

func (h *Hub) Register(conn *websocket.Conn) {
	h.register <- conn
}

func (h *Hub) Unregister(clientUUID string) {
	h.unregister <- clientUUID
}

func (h *Hub) Start() {
	go func() {
		for {
			select {
			case conn := <-h.register:
				clientUUID := uuid.Generate()
				client := NewClient(clientUUID, h, conn)
				go client.readPump()
				go client.writePump()
				h.clients[clientUUID] = client
				h.l.Debug("registered client in hub", logClientUUID(clientUUID))

			case clientUUID := <-h.unregister:
				if client, ok := h.clients[clientUUID]; ok {
					close(client.send)
					delete(h.clients, clientUUID)
					h.l.Debug("unregistered client from hub", logClientUUID(clientUUID))
				}
			case notification := <-h.broadcast:
				notificationJSON := notification.ToJSON()
				h.l.Debug("broadcasting notification to clients", zap.String("notificationJSON", notificationJSON))
				notificationBytes := []byte(notificationJSON)
				for clientUUID, client := range h.clients {
					select {
					case client.send <- notificationBytes:
						h.l.Debug("sent notification to client", logClientUUID(clientUUID))
					default:
						close(client.send)
						delete(h.clients, clientUUID)
						h.l.Debug("closed connection for client", logClientUUID(clientUUID))
					}
				}
			}
		}
	}()
}

func logClientUUID(uuid string) zap.Field {
	return zap.String("clientUUID", uuid)
}
