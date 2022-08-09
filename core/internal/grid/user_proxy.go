package grid

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/fp"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"go.uber.org/zap"
)

type UserProxy interface {
	SubscribeUserToTag(tag *domain.Tag)
	RegisterClient(interface{})
	GetLastNotifications(n int) []model.ClientNotification
	GetNotifications(from, to int) []model.ClientNotification
	GetAllNotifications() []model.ClientNotification
	GetNotificationCount() int
}

type UserProxyImpl struct {
	user        *domain.User
	repo        domain.NotificationRepository
	broadcaster domain.NotificationBroadcaster
	l           *zap.Logger
}

// @impl
var _ UserProxy = (*UserProxyImpl)(nil)

func NewUserProxy(username string, repo domain.NotificationRepository, broadcaster domain.NotificationBroadcaster) *UserProxyImpl {
	userID := uuid.Generate()
	user := domain.NewUser(userID, username, repo, broadcaster)
	return &UserProxyImpl{
		user:        user,
		repo:        repo,
		broadcaster: broadcaster,
		l: log.For("impl").Named("grid").Named("user_proxy").With(
			zap.String("userID", userID),
			zap.String("username", username),
		),
	}
}

func (p *UserProxyImpl) SubscribeUserToTag(tag *domain.Tag) {
	p.user.SubscribeToTag(tag)
	p.l.Debug("subscribed user to tag", zap.String("tagName", tag.Name))
}

func (p *UserProxyImpl) RegisterClient(conn interface{}) {
	p.broadcaster.RegisterClient(conn)
}

func (p *UserProxyImpl) GetLastNotifications(n int) []model.ClientNotification {
	domainNotifications := p.user.GetLastNotifications(n)
	return fp.Map(model.ClientNotificationFromDomain, domainNotifications)
}

func (p *UserProxyImpl) GetNotifications(from, to int) []model.ClientNotification {
	domainNotifications := p.user.GetNotifications(from, to)
	return fp.Map(model.ClientNotificationFromDomain, domainNotifications)
}

func (p *UserProxyImpl) GetAllNotifications() []model.ClientNotification {
	return p.GetNotifications(0, p.GetNotificationCount())
}

func (p *UserProxyImpl) GetNotificationCount() int {
	return p.user.GetNotificationCount()
}
