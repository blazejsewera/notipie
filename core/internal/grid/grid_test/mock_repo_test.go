package grid_test

import "github.com/blazejsewera/notipie/core/internal/domain"

type mockRepository struct{}

var _ domain.NotificationRepository = (*mockRepository)(nil)

func newMockRepository() *mockRepository {
	return &mockRepository{}
}

func (m *mockRepository) SaveNotification(domain.Notification) {}

func (m *mockRepository) GetLastNotifications(int) []domain.Notification {
	return nil
}

func (m *mockRepository) GetNotifications(int, int) []domain.Notification {
	return nil
}

func (m *mockRepository) GetNotificationCount() int {
	return 0
}
