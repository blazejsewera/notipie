package domain

import "fmt"

type Tag struct {
	Name  string
	Users []*User
	Apps  []*App
}

func (t *Tag) RegisterUser(user *User) {
	t.Users = append(t.Users, user)
}

func (t *Tag) RegisterApp(app *App) {
	t.Apps = append(t.Apps, app)
}

func (t *Tag) Broadcast(notification Notification) error {
	if len(t.Users) == 0 {
		return fmt.Errorf(noUserWhenBroadcastErrorMessage)
	}

	for _, user := range t.Users {
		user.Receive(notification)
	}
	return nil
}

const noUserWhenBroadcastErrorMessage = "no users to broadcast to"
