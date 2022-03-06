package grid

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/impl/model"
	"github.com/blazejsewera/notipie/core/internal/impl/net/ws"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/uuid"
	"go.uber.org/zap"
)

type Grid interface {
	Start()
	AddUser(username string)
	GetRootTag() *domain.Tag
	GetAppNotificationChan() chan model.AppNotification
	GetAppIDChan() chan string
	GetUserProxy(userID string) (UserProxy, error)
}

type GridImpl struct {
	rootTag             *domain.Tag
	apps                map[string]AppProxy
	users               map[string]UserProxy
	appNotificationChan chan model.AppNotification
	appIDChan           chan string
	clientHubFactory    ws.ClientHubFactory
	l                   *zap.Logger
}

func NewGrid(clientHubFactory ws.ClientHubFactory) *GridImpl {
	return &GridImpl{
		rootTag:             domain.NewTag(RootTagName),
		apps:                make(map[string]AppProxy),
		users:               make(map[string]UserProxy),
		appNotificationChan: make(chan model.AppNotification),
		appIDChan:           make(chan string),
		clientHubFactory:    clientHubFactory,
		l:                   log.For("impl").Named("grid").Named("grid"),
	}
}

func (g *GridImpl) Start() {
	g.createAndStartRootUser()
	g.rootTag.Start()
	go func() {
		for {
			g.receive(<-g.appNotificationChan)
		}
	}()

	g.l.Debug("started grid")
}

func (g *GridImpl) createAndStartRootUser() {
	g.AddUser(RootUsername)
	up, _ := g.GetUserProxy(RootUsername)
	up.SubscribeUserToTag(g.GetRootTag())

	g.l.Debug("created and started root user")
}

func (g *GridImpl) AddUser(username string) {
	up := NewUserProxy(username, g.clientHubFactory.GetClientHub())
	g.users[username] = up
	up.Start()
}

func (g *GridImpl) receive(appNotification model.AppNotification) {
	g.l.Debug("received appNotification", zap.Reflect("appNotification", appNotification))
	ap := g.getOrCreateAppProxy(appNotification)
	g.sendBackAppID(ap)
	ap.Receive(appNotification)
}

func (g *GridImpl) sendBackAppID(ap AppProxy) {
	g.l.Debug("sending back app id", zap.String("appID", ap.GetAppID()))
	g.appIDChan <- ap.GetAppID()
	g.l.Debug("sent back app id")
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
	app.Start()

	return app
}

func (g *GridImpl) GetRootTag() *domain.Tag {
	return g.rootTag
}

func (g *GridImpl) GetAppNotificationChan() chan model.AppNotification {
	return g.appNotificationChan
}

func (g *GridImpl) GetAppIDChan() chan string {
	return g.appIDChan
}

func (g *GridImpl) GetUserProxy(username string) (UserProxy, error) {
	userProxy, ok := g.users[username]
	if !ok {
		return nil, fmt.Errorf("client was not found in grid")
	}
	return userProxy, nil
}

const (
	RootUsername = "root"
	RootTagName  = "root"
)
