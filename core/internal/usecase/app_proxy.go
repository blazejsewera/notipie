package usecase

import "github.com/jazzsewera/notipie/core/internal/domain"

type AppProxy struct {
	app domain.App
}

func (p *AppProxy) OnReceive(notification domain.Notification) error {
	return p.app.Send(notification)
}
