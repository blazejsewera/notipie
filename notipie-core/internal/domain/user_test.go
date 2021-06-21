package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	User             *User
	Handler          *MockUserHandler
	TestNotification Notification
	TestError        UserHandlerError
}

func (s *UserTestSuite) SetupTest() {
	s.Handler = new(MockUserHandler)
	s.User = &User{handler: s.Handler}
	s.TestNotification = getTestNotification()
	s.TestError = getTestUserHandlerError()
}

func (s *UserTestSuite) TestReceive() {
	s.User.Receive(Application{}, s.TestNotification)
	s.Equal(s.TestNotification, s.Handler.HandledNotification)
}

func TestUserTestSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
