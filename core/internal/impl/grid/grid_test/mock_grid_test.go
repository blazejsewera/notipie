package grid_test

import "github.com/jazzsewera/notipie/core/internal/domain"

type MockGrid struct {
	RootTag              *domain.Tag
	MockApp              *domain.App
	MockUser             *domain.User
	ExpectedNotification domain.Notification
	MockUserRepository   *MockUserRepository
}

type mockCommandHandler struct{}

func (h mockCommandHandler) HandleCommand(domain.Command) {}

type MockUserRepository struct {
	Done             chan struct{}
	lastNotification domain.Notification
	count            int
}

func (m *MockUserRepository) SaveNotification(notification domain.Notification) {
	m.Done <- struct{}{}
	m.lastNotification = notification
	m.count++
}

func (m *MockUserRepository) GetLastNotifications(int) []domain.Notification {
	return []domain.Notification{m.lastNotification}
}

func (m *MockUserRepository) GetNotifications(int, int) []domain.Notification {
	return []domain.Notification{m.lastNotification}
}

func (m *MockUserRepository) GetNotificationCount() int {
	return m.count
}

func NewMockGrid() *MockGrid {
	rootTag := &domain.Tag{Name: "root"}
	handler := mockCommandHandler{}
	mockApp := domain.NewApp("appId", "TestApp", "", "", handler)
	repo := &MockUserRepository{Done: make(chan struct{})}
	mockUser := domain.NewUser("userId", "TestUser", repo)
	return &MockGrid{
		RootTag:            rootTag,
		MockApp:            mockApp,
		MockUser:           mockUser,
		MockUserRepository: repo,
	}
}

func (g *MockGrid) Start() {
	g.MockApp.AddTag(g.RootTag)
	g.MockUser.SubscribeToTag(g.RootTag)
	g.RootTag.Listen()
	g.MockApp.Start()
	g.MockUser.Listen()
}

func (g *MockGrid) GetRootTag() *domain.Tag {
	return g.RootTag
}
