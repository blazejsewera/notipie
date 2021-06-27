package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type AppSuite struct {
	suite.Suite
	App              *App
	TestNotification Notification
	// TODO: Add App Notification Repo
}

func (s *AppSuite) SetupTest() {
	s.App = &App{}
	s.TestNotification = getTestNotification()
}

func (s *AppSuite) TestSendNotification() {
	s.App.Send(s.TestNotification)
	// TODO: Add repository content check
}

func TestApp(t *testing.T) {
	suite.Run(t, new(AppSuite))
}
