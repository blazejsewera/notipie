package usecase

import "github.com/jazzsewera/notipie/notipie-core/internal/domain"

type AppProxy struct {
	app domain.App
}

func (p *AppProxy) OnReceive(notification domain.Notification) error {
	p.app.Send(notification)
	return nil
}
