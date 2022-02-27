package grid

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/internal/impl/net/ws"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
)

type Grid interface {
	Start()
	AddUser(username string)
	GetRootTag() *domain.Tag
	GetAppNotificationChan() chan model.AppNotification
	GetUserProxy(userID string) (UserProxy, error)
}

type GridImpl struct {
	rootTag             *domain.Tag
	apps                map[string]AppProxy
	users               map[string]UserProxy
	appNotificationChan chan model.AppNotification
	clientHubFactory    ws.ClientHubFactory
}

func NewGrid(clientHubFactory ws.ClientHubFactory) *GridImpl {
	return &GridImpl{
		rootTag:             domain.NewTag(RootTagName),
		apps:                make(map[string]AppProxy),
		users:               make(map[string]UserProxy),
		appNotificationChan: make(chan model.AppNotification),
		clientHubFactory:    clientHubFactory,
	}
}

func (g *GridImpl) Start() {
	g.createAndStartRootUser()
	g.rootTag.Listen()
	go func() {
		for {
			g.receive(<-g.appNotificationChan)
		}
	}()
}

func (g *GridImpl) createAndStartRootUser() {
	g.AddUser(RootUsername)
	up, _ := g.GetUserProxy(RootUsername)
	up.SubscribeUserToTag(g.GetRootTag())
}

func (g *GridImpl) AddUser(username string) {
	up := NewUserProxy(username, g.clientHubFactory.GetClientHub())
	g.users[username] = up
	up.Start()
}

func (g *GridImpl) receive(appNotification model.AppNotification) {
	ap := g.getOrCreateAppProxy(appNotification)
	ap.Receive(appNotification)
}

func (g *GridImpl) getOrCreateAppProxy(n model.AppNotification) AppProxy {
	appID := n.AppID

	ap, ok := g.apps[appID]
	if ok {
		return ap
	}

	return g.createAppProxy(n)

}

func (g *GridImpl) createAppProxy(n model.AppNotification) AppProxy {
	app := g.createAndInitDomainApp(n)
	return NewAppProxy(app)
}

func (g *GridImpl) createAndInitDomainApp(n model.AppNotification) *domain.App {
	appID := uuid.Generate()

	app := &domain.App{
		ID:      appID,
		Name:    n.AppName,
		IconURI: n.AppImgURI,
	} // TODO: implement command handler, use NewApp()

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
