package grid

import "github.com/jazzsewera/notipie/core/internal/impl/model"

type UserProxy interface {
	GetClientNotificationChan() chan model.ClientNotification
}
