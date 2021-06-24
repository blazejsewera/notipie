package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	User                          *User
	Repo                          *MockUserNotificationRepository
	Handler                       *MockUserHandler
	TestNotification              Notification
	TestErrNotification           Notification
	TestUserHandlerError          UserHandlerError
	TestUserNotificationRepoError UserNotificationRepositoryError
}

func (s *UserTestSuite) SetupTest() {
	s.Handler = new(MockUserHandler)
	s.User = &User{handler: s.Handler}
	s.TestNotification = getTestNotification()
	s.TestErrNotification = getTestErrNotification()
	s.TestUserHandlerError = getTestUserHandlerError()
	s.TestUserNotificationRepoError = getTestUserNotificationRepoError()
}

func (s *UserTestSuite) TestReceive() {
	s.User.Receive(Application{}, s.TestNotification)
	s.Equal(s.TestNotification, s.Handler.HandledNotification)
}

func (s *UserTestSuite) TestRepository() {
	s.Run("save notification", func() {
		// when
		_ = s.User.repo.SaveNotification(s.TestNotification)

		// then
		// TODO: Refactor test to use a field, so that only one method at a time is tested
		notifications, _ := s.User.repo.GetNotifications()
		s.ElementsMatch([...]Notification{s.TestNotification}, notifications)
	})

	s.Run("save err notification", func() {
		// when
		err := s.User.repo.SaveNotification(s.TestErrNotification)

		// then
		if err != nil {
			s.User.repo.HandleError(err)
		}
		s.Equal(s.TestUserNotificationRepoError, s.Repo.Err)
	})
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
