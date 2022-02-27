package grid

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/internal/impl/persistence"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
)

type UserProxy interface {
	GetClientHub() ws.ClientHub
	SubscribeUserToTag(tag *domain.Tag)
}

type UserProxyImpl struct {
	user *domain.User
	repo persistence.RealtimeNotificationRepo
	hub  ws.ClientHub
}

func NewUserProxy(username string, hub ws.ClientHub) *UserProxyImpl {
	repo := persistence.NewMemRealtimeNotificationRepository()
	userID := uuid.Generate()
	user := domain.NewUser(userID, username, repo)
	return &UserProxyImpl{
		user: user,
		repo: repo,
		hub:  hub,
	}
}

func (p *UserProxyImpl) Start() {
	p.user.Listen()
	p.hub.Run()
	go func() {
		p.hub.GetBroadcastChan() <- clientNotificationOf(<-p.repo.GetNotificationChan())
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

func (p *UserProxyImpl) GetClientHub() ws.ClientHub {
	return p.hub
}

func (p *UserProxyImpl) SubscribeUserToTag(tag *domain.Tag) {
	p.user.SubscribeToTag(tag)
}
