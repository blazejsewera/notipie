package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ApplicationTestSuite struct {
	suite.Suite
	App                 *Application
	Handler             *MockAppHandler
	TestNotification    Notification
	TestErrNotification Notification
	TestError           AppHandlerError
}

func (s *ApplicationTestSuite) SetupTest() {
	s.Handler = new(MockAppHandler)
	s.App = &Application{handler: s.Handler}
	s.TestNotification = getTestNotification()
	s.TestErrNotification = getTestErrNotification()
	s.TestError = getTestAppHandlerError()
}

func (s ApplicationTestSuite) TestSendNotification() {
	s.App.Send(s.TestNotification)
	s.Equal(s.TestNotification, s.Handler.HandledNotification)
}

func (s ApplicationTestSuite) TestFailToSendNotification() {
	s.App.Send(s.TestErrNotification)
	s.Equal(s.TestError, s.Handler.Err)
}

func (s ApplicationTestSuite) TestEquals() {
	app1 := Application{ID: "ID1"}
	app1Copy := Application{ID: "ID1"}
	app2 := Application{ID: "ID2"}
	s.True(app1.Equals(app1Copy))
	s.False(app1.Equals(app2))
}

func TestApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}
