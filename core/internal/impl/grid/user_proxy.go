package grid

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/impl/broadcast"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/internal/impl/persistence"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"go.uber.org/zap"
)

type UserProxy interface {
	GetHub() ws.Hub
	SubscribeUserToTag(tag *domain.Tag)
}

type UserProxyImpl struct {
	user        *domain.User
	repo        *persistence.MemRealtimeNotificationRepository
	broadcaster *broadcast.WebSocketNotificationBroadcaster
	hub         ws.Hub
	l           *zap.Logger
}

// UserProxyImpl implements interfaces below
var _ UserProxy = (*UserProxyImpl)(nil)
var _ util.Starter = (*UserProxyImpl)(nil)

func NewUserProxy(username string, hub ws.Hub) *UserProxyImpl {
	repo := persistence.NewMemRealtimeNotificationRepository()
	broadcaster := broadcast.NewWebSocketNotificationBroadcaster(hub)
	userID := uuid.Generate()
	user := domain.NewUser(userID, username, repo, broadcaster)
	return &UserProxyImpl{
		user:        user,
		repo:        repo,
		broadcaster: broadcaster,
		hub:         hub,
		l: log.For("impl").Named("grid").Named("user_proxy").With(
			zap.String("userID", userID),
			zap.String("username", username),
		),
	}
}

func (p *UserProxyImpl) Start() {
	p.user.Start()
	p.broadcaster.Start()
	p.l.Debug("started user proxy")
}

func (p *UserProxyImpl) GetHub() ws.Hub {
	return p.hub
}

func (p *UserProxyImpl) SubscribeUserToTag(tag *domain.Tag) {
	p.user.SubscribeToTag(tag)
	p.l.Debug("subscribed user to tag", zap.String("tagName", tag.Name))
}
