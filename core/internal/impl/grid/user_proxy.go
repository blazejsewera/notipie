package grid

import (
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/internal/impl/net/ws"
	"github.com/jazzsewera/notipie/core/internal/impl/persistence"
	"github.com/jazzsewera/notipie/core/pkg/lib/timeformat"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
)

type UserProxy interface {
	GetClientNotificationChan() chan model.ClientNotification
	GetClientHub() ws.ClientHub
	SubscribeUserToTag(tag *domain.Tag)
}

type UserProxyImpl struct {
	user                   *domain.User
	repo                   persistence.RealtimeNotificationRepo
	hub                    ws.ClientHub
	clientNotificationChan chan model.ClientNotification
}

func NewUserProxy(username string) *UserProxyImpl {
	repo := persistence.NewMemRealtimeNotificationRepository()
	userID := uuid.Generate()
	user := domain.NewUser(userID, username, repo)
	hub := ws.NewHub()
	return &UserProxyImpl{
		user:                   user,
		repo:                   repo,
		hub:                    hub,
		clientNotificationChan: make(chan model.ClientNotification),
	}
}

func (p *UserProxyImpl) Start() {
	p.user.Listen()
	go func() {
		p.clientNotificationChan <- clientNotificationOf(<-p.repo.GetNotificationChan())
	}()
}

func clientNotificationOf(n domain.Notification) model.ClientNotification {
	timestamp := n.Timestamp.Format(timeformat.RFC3339Milli)
	return model.ClientNotification{
		HashableNetNotification: model.HashableNetNotification{
			AppName:    n.App.Name,
			AppID:      n.App.ID,
			AppImgURI:  n.App.IconURI,
			Title:      n.Title,
			Subtitle:   n.Subtitle,
			Body:       n.Body,
			ExtURI:     n.ExtURI,
			ReadURI:    n.ReadURI,
			ArchiveURI: n.ArchiveURI,
		},
		ID:        n.ID,
		Timestamp: timestamp,
		Read:      false,
	} // TODO: implement urgency
}

func (p *UserProxyImpl) GetClientNotificationChan() chan model.ClientNotification {
	return p.clientNotificationChan
}

func (p *UserProxyImpl) GetClientHub() ws.ClientHub {
	return p.hub
}

func (p *UserProxyImpl) SubscribeUserToTag(tag *domain.Tag) {
	p.user.SubscribeToTag(tag)
}
