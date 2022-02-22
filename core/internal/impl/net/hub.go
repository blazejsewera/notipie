package net

import (
	"encoding/json"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
)

type ClientHub interface {
	GetBroadcastChan() chan model.ClientNotification
}

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan model.ClientNotification
	register   chan *Client
	unregister chan *Client
	l          *zap.Logger
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan model.ClientNotification),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		l:          log.For("net").Named("hub"),
	}
}

func (h *Hub) GetBroadcastChan() chan model.ClientNotification {
	return h.broadcast
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case notification := <-h.broadcast:
			notificationBytes, err := json.Marshal(notification)
			if err != nil {
				h.l.Warn("could not serialize notification", zap.Error(err))
			}
			for client := range h.clients {
				select {
				case client.send <- notificationBytes:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
