package impl

import (
	"encoding/json"
	"log"

	"github.com/jazzsewera/notipie/core/internal/domain"
)

type Hub struct {
	user       *domain.User
	clients    map[*Client]bool
	broadcast  chan domain.Notification
	register   chan *Client
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan domain.Notification),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
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
			notification_bytes, err := json.Marshal(notification)
			if err != nil {
				log.Printf("could not serialize notification")
			}
			for client := range h.clients {
				select {
				case client.send <- notification_bytes:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}
