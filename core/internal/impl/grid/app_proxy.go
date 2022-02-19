package grid

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	timeFormat "github.com/jazzsewera/notipie/core/pkg/lib/time"
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
	app := p.getOrCreateApp(netNotification.AppName)
	timestamp, err := time.Parse(timeFormat.RFC3339Milli, netNotification.Timestamp)
	if err != nil {
		return domain.Notification{}, err
	}

	notification := domain.Notification{
		ID:        netNotification.ID,
		App:       app,
		Timestamp: timestamp,
		Title:     netNotification.Title,
		Body:      netNotification.Body,
		Urgency:   domain.Medium, // TODO: Implement urgency
	}

	return notification, nil
}

func (p *AppProxy) getOrCreateApp(appName string) *domain.App {
	app, ok := p.apps[appName]
	if ok {
		return app
	}

	return p.createAndInitApp(appName)
}

func (p *AppProxy) createAndInitApp(appName string) *domain.App {
	app := &domain.App{
		ID:   uuid.Generate(),
		Name: appName,
		// TODO: Implement small and big icon img
	}

	p.apps[appName] = app
	app.AddTag(p.grid.GetRootTag())
	app.Start()

	return app
}
