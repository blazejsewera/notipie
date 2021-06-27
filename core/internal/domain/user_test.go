package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserSuite struct {
	suite.Suite
	User             *User
	Repo             *MockUserNotificationRepository
	TestNotification Notification
	Tag              Tag
}

func (s *UserSuite) SetupTest() {
	s.User = &User{}
	s.TestNotification = getTestNotification()
}

func (s *UserSuite) TestReceive() {
	s.User.Receive(s.TestNotification)
	// TODO: Check user repo content
}

func (s *UserSuite) TestRepository() {
	s.Run("save notification", func() {
		// when
		s.User.repo.SaveNotification(s.TestNotification)

		// then
		// TODO: Refactor test to use a field, so that only one method at a time is tested
		notifications := s.User.repo.GetAllNotifications()
		s.ElementsMatch([...]Notification{s.TestNotification}, notifications)
	})
}

func (s *UserSuite) TestSubscribeToTag() {
	s.User.SubscribeTo(&s.Tag)
	s.ElementsMatch([...]*Tag{&s.Tag}, s.User.tags)
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserSuite))
}
