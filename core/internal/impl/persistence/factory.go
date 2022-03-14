package persistence

import "github.com/blazejsewera/notipie/core/internal/domain"

type RepositoryFactoryFunc func() domain.NotificationRepository

func (f RepositoryFactoryFunc) GetRepository() domain.NotificationRepository {
	return f()
}

var MemRepositoryFactory = RepositoryFactoryFunc(func() domain.NotificationRepository {
	return NewMemRepository()
})
