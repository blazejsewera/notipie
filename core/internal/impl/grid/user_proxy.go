package grid

import "github.com/jazzsewera/notipie/core/internal/impl/model"

type UserProxy interface {
	GetClientNotificationChan() chan model.ClientNotification
}

type UserProxyImpl struct {
	ClientNotificationChan chan model.ClientNotification
}

func NewUserProxy() *UserProxyImpl {
	return &UserProxyImpl{ClientNotificationChan: make(chan model.ClientNotification)}
}

func (p *UserProxyImpl) GetClientNotificationChan() chan model.ClientNotification {
	return p.ClientNotificationChan
}
