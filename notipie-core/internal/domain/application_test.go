package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ApplicationTestSuite struct {
	suite.Suite
	App              *Application
	ErrApp           *Application
	Handler          *MockHandler
	ErrHandler       *ErroredHandler
	TestNotification Notification
	TestError        HandlerError
}

func (s *ApplicationTestSuite) SetupTest() {
	s.Handler = new(MockHandler)
	s.ErrHandler = new(ErroredHandler)
	s.App = &Application{handler: s.Handler}
	s.ErrApp = &Application{handler: s.ErrHandler}
	s.TestNotification = getTestNotification()
	s.TestError = getTestHandlerError()
}

func (s ApplicationTestSuite) TestSendNotification() {
	s.App.Send(s.TestNotification)
	s.Equal(s.TestNotification, s.Handler.HandledNotification)
}

func (s ApplicationTestSuite) TestFailToSendNotification() {
	s.ErrApp.Send(s.TestNotification)
	s.Equal(s.TestError, s.ErrHandler.Err)
}

func (s ApplicationTestSuite) TestEquals() {
	app1 := &Application{ID: "ID1"}
	app1Copy := &Application{ID: "ID1"}
	app2 := &Application{ID: "ID2"}
	s.True(app1.Equals(app1Copy))
	s.False(app1.Equals(app2))
}

func TestApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationTestSuite))
}
