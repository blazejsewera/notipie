package ws

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/pkg/lib/log"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
	"go.uber.org/zap"
)

type ClientHub interface {
	GetBroadcastChan() chan model.ClientNotification
	GetRegisterChan() chan *websocket.Conn
	GetUnregisterChan() chan string
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

func (h *Hub) Run() {
	for {
		select {
		case conn := <-h.register:
			clientUUID := uuid.Generate()
			client := NewClient(clientUUID, h, conn)
			go client.readPump()
			go client.writePump()
			h.clients[clientUUID] = client

		case clientUUID := <-h.unregister:
			if client, ok := h.clients[clientUUID]; ok {
				close(client.send)
				delete(h.clients, clientUUID)
			}
		case notification := <-h.broadcast:
			notificationBytes, err := json.Marshal(notification)
			if err != nil {
				h.l.Warn("could not serialize notification", zap.Error(err))
			}
			for clientUUID, client := range h.clients {
				select {
				case client.send <- notificationBytes:
				default:
					close(client.send)
					delete(h.clients, clientUUID)
				}
			}
		}
	}
}
