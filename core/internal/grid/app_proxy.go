package grid

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/model"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"github.com/blazejsewera/notipie/core/pkg/lib/timeformat"
	"go.uber.org/zap"
	"time"
)

type AppProxy interface {
	GetAppID() string
	receive(appNotification model.AppNotification)
}

type AppProxyImpl struct {
	app *domain.App
	l   *zap.Logger
}

//@impl
var _ AppProxy = (*AppProxyImpl)(nil)

func NewAppProxy(app *domain.App) *AppProxyImpl {
	return &AppProxyImpl{app: app, l: log.For("impl").Named("grid").Named("app_proxy")}
}

func (p *AppProxyImpl) receive(appNotification model.AppNotification) {
	p.l.Debug("received appNotification", zap.Reflect("appNotification", appNotification))
	notification, err := p.notificationOf(appNotification)
	if err != nil {
		p.l.Error("error when converting a notification", zap.Error(err))
		return
	}

	app := notification.App
	err = app.Send(notification)

	if err != nil {
		p.l.Error("error when sending a notification", zap.Error(err))
		return
	}
	p.l.Debug("sent domain notification with app", zap.String("appID", app.ID), zap.String("appName", app.Name))
}

func (p *AppProxyImpl) notificationOf(appNotification model.AppNotification) (domain.Notification, error) {
	appNotification.AppID = p.app.ID
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

func (p *AppProxyImpl) GetAppID() string {
	return p.app.ID
}
