package ws

import (
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Hub interface {
	Broadcast(notification model.ClientNotification)
	Register(conn *websocket.Conn)
	Unregister(clientUUID string)
}

type HubImpl struct {
	clientFactory ClientFactory
	clients       map[string]Client
	broadcast     chan model.ClientNotification
	register      chan *websocket.Conn
	unregister    chan string
	l             *zap.Logger
}

//@impl
var _ Hub = (*HubImpl)(nil)

func NewHubImpl(clientFactory ClientFactory) *HubImpl {
	h := &HubImpl{
		clientFactory: clientFactory,
		broadcast:     make(chan model.ClientNotification),
		register:      make(chan *websocket.Conn),
		unregister:    make(chan string),
		clients:       make(map[string]Client),
		l:             log.For("impl").Named("net").Named("hub"),
	}
	clientFactory.SetHub(h)
	h.start()

	return h
}

func (h *HubImpl) start() {
	go h.handleChannels()
}

func (h *HubImpl) Broadcast(notification model.ClientNotification) {
	h.broadcast <- notification
}

func (h *HubImpl) Register(conn *websocket.Conn) {
	h.register <- conn
}

func (h *HubImpl) Unregister(clientUUID string) {
	h.unregister <- clientUUID
}

func (h *HubImpl) handleChannels() {
	for {
		select {
		case conn := <-h.register:
			h.registerClient(conn)
		case clientUUID := <-h.unregister:
			h.unregisterClient(clientUUID)
		case notification := <-h.broadcast:
			h.broadcastNotification(notification)
		}
	}
}

func (h *HubImpl) registerClient(conn *websocket.Conn) {
	client := NewClient(h, conn)
	client.Start()
	h.clients[client.UUID()] = client
	h.l.Debug("registered client in hub", logClientUUID(client.UUID()))
}

func (h *HubImpl) unregisterClient(clientUUID string) {
	if _, ok := h.clients[clientUUID]; ok {
		delete(h.clients, clientUUID)
		h.l.Debug("unregistered client from hub", logClientUUID(clientUUID))
	}
}

func (h *HubImpl) broadcastNotification(notification model.ClientNotification) {
	notificationBytes, err := notification.ToJSON()
	if err != nil {
		h.l.Error("broadcast notification", zap.Error(err))
		return
	}
	for _, client := range h.clients {
		client.Broadcast(notificationBytes)
	}
}

func logClientUUID(uuid string) zap.Field {
	return zap.String("clientUUID", uuid)
}
