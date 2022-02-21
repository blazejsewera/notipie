package grid

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/pkg/lib/timeformat"
	"github.com/jazzsewera/notipie/core/pkg/lib/uuid"
	"time"
)

type AppProxy struct {
	NetNotificationChan chan model.NetNotification
	grid                Grid
	apps                map[string]*domain.App
}

func NewAppProxy(grid Grid) *AppProxy {
	return &AppProxy{grid: grid, apps: make(map[string]*domain.App)}
}

func (p *AppProxy) Listen() {
	if p.NetNotificationChan == nil {
		p.NetNotificationChan = make(chan model.NetNotification)
	}
	go func() {
		for {
			p.Receive(<-p.NetNotificationChan)
		}
	}()
}

func (p *AppProxy) GetAppCount() int {
	return len(p.apps)
}

func (p *AppProxy) Receive(netNotification model.NetNotification) {
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

func (p *AppProxy) notificationOf(netNotification model.NetNotification) (domain.Notification, error) {
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

func (p *AppProxy) getOrCreateApp(n model.NetNotification) *domain.App {
	appID := n.AppID

	app, ok := p.apps[appID]
	if ok {
		return app
	}

	return p.createAndInitApp(n)

}

func (p *AppProxy) createAndInitApp(n model.NetNotification) *domain.App {
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
