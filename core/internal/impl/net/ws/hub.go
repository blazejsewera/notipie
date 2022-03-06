package ws

import (
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ClientHub interface {
	GetBroadcastChan() chan model.ClientNotification
	GetRegisterChan() chan *websocket.Conn
	GetUnregisterChan() chan string
	Start()
}

type ClientHubFactory interface {
	GetClientHub() ClientHub
}

type DefaultClientHubFactory struct{}

func (f DefaultClientHubFactory) GetClientHub() ClientHub {
	return NewHub()
}

type Hub struct {
	clients    map[string]*Client
	broadcast  chan model.ClientNotification
	register   chan *websocket.Conn
	unregister chan string
	l          *zap.Logger
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan model.ClientNotification),
		register:   make(chan *websocket.Conn),
		unregister: make(chan string),
		clients:    make(map[string]*Client),
		l:          log.For("impl").Named("net").Named("hub"),
	}
}

func (h *Hub) GetBroadcastChan() chan model.ClientNotification {
	return h.broadcast
}

func (h *Hub) GetRegisterChan() chan *websocket.Conn {
	return h.register
}

func (h *Hub) GetUnregisterChan() chan string {
	return h.unregister
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
