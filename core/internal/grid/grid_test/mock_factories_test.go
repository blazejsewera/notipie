package grid_test

import (
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/blazejsewera/notipie/core/internal/impl/broadcast"
	"github.com/blazejsewera/notipie/core/internal/impl/persistence"
)

var mockBroadcasterFactory = broadcast.BroadcasterFactoryFunc(func() domain.NotificationBroadcaster {
	return mockBroadcasterInstance
})

var mockRepositoryFactory = persistence.RepositoryFactoryFunc(func() domain.NotificationRepository {
	return newMockRepository()
})
