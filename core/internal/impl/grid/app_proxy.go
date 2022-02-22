package grid

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/pkg/lib/timeformat"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
	"time"
)

type AppProxy interface {
	GetAppNotificationChan() chan model.AppNotification
	GetAppCount() int
}

type AppProxyImpl struct {
	AppNotificationChan chan model.AppNotification
	grid                Grid
	apps                map[string]*domain.App
}

func NewAppProxy(grid Grid) *AppProxyImpl {
	return &AppProxyImpl{grid: grid, apps: make(map[string]*domain.App)}
}

func (p *AppProxyImpl) Listen() {
	if p.AppNotificationChan == nil {
		p.AppNotificationChan = make(chan model.AppNotification)
	}
	go func() {
		for {
			p.Receive(<-p.AppNotificationChan)
		}
	}()
}

func (p *AppProxyImpl) GetAppNotificationChan() chan model.AppNotification {
	return p.AppNotificationChan
}

func (p *AppProxyImpl) GetAppCount() int {
	return len(p.apps)
}

func (p *AppProxyImpl) Receive(netNotification model.AppNotification) {
	notification, err := p.notificationOf(netNotification)
	if err != nil {
		fmt.Printf("error when converting a notification: %s", err)
		return
	}

	app := notification.App
	err = app.Send(notification)

	if err != nil {
		fmt.Printf("error when sending a notification: %s", err)
		return
	}
}

func (p *AppProxyImpl) notificationOf(netNotification model.AppNotification) (domain.Notification, error) {
	nn := model.AddIDTo(netNotification)
	app := p.getOrCreateApp(nn)
	timestamp, err := time.Parse(timeformat.RFC3339Milli, nn.Timestamp)
	if err != nil {
		return domain.Notification{}, err
	}

	notification := domain.Notification{
		ID:        nn.ID,
		App:       app,
		Timestamp: timestamp,
		Title:     nn.Title,
		Body:      nn.Body,
		Urgency:   domain.Medium, // TODO: Implement urgency
	}

	return notification, nil
}

func (p *AppProxyImpl) getOrCreateApp(n model.AppNotification) *domain.App {
	appID := n.AppID

	app, ok := p.apps[appID]
	if ok {
		return app
	}

	return p.createAndInitApp(n)

}

func (p *AppProxyImpl) createAndInitApp(n model.AppNotification) *domain.App {
	appID := uuid.Generate()

	app := &domain.App{
		ID:      appID,
		Name:    n.AppName,
		IconURI: n.AppImgURI,
	}

	p.apps[appID] = app
	app.AddTag(p.grid.GetRootTag())
	app.Start()

	return app
}
