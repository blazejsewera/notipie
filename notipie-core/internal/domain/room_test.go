package domain

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type RoomTestSuite struct {
	suite.Suite
	RoomU0A0             Room
	RoomU01A1            Room
	RoomU012A2           Room
	RoomU2A23            Room
	Users                []User
	UserHandlers         []MockUserHandler
	Apps                 []Application
	AppHandlers          []MockAppHandler
	TestErrNotification  Notification
	TestAppHandlerError  AppHandlerError
	TestUserHandlerError UserHandlerError
}

func (s *RoomTestSuite) SetupTest() {
	s.UserHandlers = make([]MockUserHandler, 3)
	s.AppHandlers = make([]MockAppHandler, 4)
	s.Users = []User{
		{ID: "0", Username: "User0", handler: &s.UserHandlers[0]},
		{ID: "1", Username: "User1", handler: &s.UserHandlers[1]},
		{ID: "2", Username: "User2", handler: &s.UserHandlers[2]},
	}
	s.Apps = []Application{
		{ID: "0", Name: "App0", SmallIconURL: "sApp0", BigIconURL: "bApp0", handler: &s.AppHandlers[0]},
		{ID: "1", Name: "App1", SmallIconURL: "sApp1", BigIconURL: "bApp1", handler: &s.AppHandlers[1]},
		{ID: "2", Name: "App2", SmallIconURL: "sApp2", BigIconURL: "bApp2", handler: &s.AppHandlers[2]},
		{ID: "3", Name: "App3", SmallIconURL: "sApp3", BigIconURL: "bApp3", handler: &s.AppHandlers[3]},
	}
	s.RoomU0A0 = Room{
		Users: s.Users[0:0], // s.Users[0]
		Apps:  s.Apps[0:0],  // s.Apps[0]
	}
	s.RoomU01A1 = Room{
		Users: s.Users[0:1], // s.Users[0], s.Users[1]
		Apps:  s.Apps[1:1],  // s.Apps[1]
	}
	s.RoomU2A23 = Room{
		Users: s.Users[2:2], // s.Users[2]
		Apps:  s.Apps[2:3],  // s.Apps[2], s.Apps[3]
	}
	s.TestErrNotification = getTestErrNotification()
	s.TestAppHandlerError = getTestAppHandlerError()
	s.TestUserHandlerError = getTestUserHandlerError()
}

func (s *RoomTestSuite) TestBroadcast() {
	s.Run("broadcast from App 0 to User 0", func() {
		// given
		notification := Notification{Title: "App 0 to User 0"}
		uh0 := s.UserHandlers[0]
		app0 := s.Apps[0]

		// when
		app0.Send(notification)

		// then
		s.True(uh0.HandledApp.Equals(app0))
		s.Equal(notification, uh0.HandledNotification)
	})

	s.Run("broadcast from App 1 to User 0 and 1", func() {
		// given
		notification := Notification{Title: "App 1 to User 0 and 1"}
		uh0 := s.UserHandlers[0]
		uh1 := s.UserHandlers[1]
		app1 := s.Apps[1]

		// when
		app1.Send(notification)

		// then
		s.True(uh0.HandledApp.Equals(app1))
		s.Equal(notification, uh0.HandledNotification)
		s.True(uh1.HandledApp.Equals(app1))
		s.Equal(notification, uh1.HandledNotification)
	})

	s.Run("broadcast from App 2 and 3 to User 2", func() {
		// given
		//notificationApp2 := Notification{Title: "App 2 to User 2"}
		//notificationApp3 := Notification{Title: "App 3 to User 2"}
	})
}

func TestRoomTestSuite(t *testing.T) {
	suite.Run(t, new(RoomTestSuite))
}
