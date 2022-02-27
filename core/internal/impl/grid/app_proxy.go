package grid

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/impl/model"
	"github.com/jazzsewera/notipie/core/pkg/lib/log"
	"github.com/jazzsewera/notipie/core/pkg/lib/timeformat"
	"go.uber.org/zap"
	"time"
)

type AppProxy interface {
	Receive(appNotification model.AppNotification)
}

type AppProxyImpl struct {
	app *domain.App
	l   *zap.Logger
}

func NewAppProxy(app *domain.App) *AppProxyImpl {
	return &AppProxyImpl{app: app, l: log.For("grid").Named("app_proxy")}
}

func (p *AppProxyImpl) Receive(appNotification model.AppNotification) {
	notification, err := p.notificationOf(appNotification)
	if err != nil {
		fmt.Printf("error when converting a notification: %s", err)
		return
	}

	app := notification.App
	err = app.Send(notification)

	if err != nil {
		p.l.Error("error when sending a notification", zap.Error(err))
		return
	}
}

func (p *AppProxyImpl) notificationOf(appNotification model.AppNotification) (domain.Notification, error) {
	an := model.AddIDTo(appNotification)
	timestamp, err := time.Parse(timeformat.RFC3339Milli, an.Timestamp)
	if err != nil {
		return domain.Notification{}, err
	}

	notification := domain.Notification{
		ID:         an.ID,
		App:        p.app,
		Timestamp:  timestamp,
		Title:      an.Title,
		Subtitle:   an.Subtitle,
		Body:       an.Body,
		Urgency:    domain.Medium, // TODO: implement urgency
		ExtURI:     an.ExtURI,
		ReadURI:    an.ReadURI,
		ArchiveURI: an.ArchiveURI,
	}

	return notification, nil
}
