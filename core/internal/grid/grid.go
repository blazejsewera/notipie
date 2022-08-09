package grid

import (
	"errors"
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/util"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"github.com/blazejsewera/notipie/core/pkg/model"
	"go.uber.org/zap"
)

type Grid interface {
	util.Starter
	AddUser(username string)
	ReceiveAppNotification(notification model.AppNotification) (appID string)
	GetUserProxy(userID string) (UserProxy, error)
	getRootTag() *domain.Tag
}

type GridImpl struct {
	rootTag      *domain.Tag
	apps         map[string]AppProxy
	users        map[string]UserProxy
	repositories RepositoryFactory
	broadcasters BroadcasterFactory
	l            *zap.Logger
}

// @impl
var _ Grid = (*GridImpl)(nil)

func NewGrid(repositoryFactory RepositoryFactory, broadcasterFactory BroadcasterFactory) *GridImpl {
	return &GridImpl{
		rootTag:      domain.NewTag(RootTagName),
		apps:         make(map[string]AppProxy),
		users:        make(map[string]UserProxy),
		repositories: repositoryFactory,
		broadcasters: broadcasterFactory,
		l:            log.For("impl").Named("grid").Named("grid"),
	}
}

func (g *GridImpl) Start() {
	g.createAndStartRootUser()
	g.l.Debug("started grid")
}

func (g *GridImpl) createAndStartRootUser() {
	g.AddUser(RootUsername)
	up, _ := g.GetUserProxy(RootUsername)
	up.SubscribeUserToTag(g.getRootTag())

	g.l.Debug("created and started root user")
}

func (g *GridImpl) AddUser(username string) {
	r := g.repositories.GetRepository()
	b := g.broadcasters.GetBroadcaster()
	up := NewUserProxy(username, r, b)
	g.users[username] = up
}

func (g *GridImpl) getRootTag() *domain.Tag {
	return g.rootTag
}

func (g *GridImpl) ReceiveAppNotification(appNotification model.AppNotification) (appID string) {
	g.l.Debug("received appNotification", zap.String("notificationID", appNotification.ID))
	ap := g.getOrCreateAppProxy(appNotification)
	ap.receive(appNotification)
	return ap.GetAppID()
}

func (g *GridImpl) getOrCreateAppProxy(n model.AppNotification) AppProxy {
	appID := n.AppID

	ap, ok := g.apps[appID]
	if ok {
		g.l.Debug("found app", zap.String("appID", appID))
		return ap
	}

	g.l.Debug("did not find app", zap.String("appID", appID))
	newAP := g.createAppProxy(n)
	newAppID := newAP.GetAppID()
	g.apps[newAppID] = newAP
	return newAP
}

func (g *GridImpl) createAppProxy(n model.AppNotification) AppProxy {
	app := g.createAndInitDomainApp(n)
	return NewAppProxy(app)
}

func (g *GridImpl) createAndInitDomainApp(n model.AppNotification) *domain.App {
	appID := uuid.Generate()
	app := domain.NewApp(appID, n.AppName, n.AppImgURI, nil) // TODO: implement command handler
	app.AddTag(g.rootTag)
	return app
}

func (g *GridImpl) GetUserProxy(username string) (UserProxy, error) {
	userProxy, ok := g.users[username]
	if !ok {
		return nil, errors.New("client was not found in grid")
	}
	return userProxy, nil
}

const (
	RootUsername = "root"
	RootTagName  = "root"
)

type RepositoryFactory interface {
	GetRepository() domain.NotificationRepository
}

type BroadcasterFactory interface {
	GetBroadcaster() domain.NotificationBroadcaster
}
