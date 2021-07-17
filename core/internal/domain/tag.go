package domain

import "fmt"

type Tag struct {
	Name             string
	Users            []*User
	Apps             []*App
	NotificationChan chan Notification
}

func (t *Tag) Listen() {
	if t.NotificationChan == nil {
		t.NotificationChan = make(chan Notification)
	}

	go func() {
		for {
			_ = t.broadcast(<-t.NotificationChan)
		}
	}()
}

func (t *Tag) broadcast(notification Notification) error {
	if len(t.Users) == 0 {
		return fmt.Errorf(noUserWhenBroadcastErrorMessage)
	}

	for _, user := range t.Users {
		user.NotificationChan <- notification
	}
	return nil
}

func (t *Tag) registerUser(user *User) {
	t.Users = append(t.Users, user)
}

func (t *Tag) registerApp(app *App) {
	t.Apps = append(t.Apps, app)
}

const (
	noUserWhenBroadcastErrorMessage     = "no users to broadcast to"
	noMatchingTagsWhenRemoveErrorFormat = "tag %q not found"
)

func removeTag(tags []*Tag, tag Tag) ([]*Tag, error) {
	var reduced []*Tag
	found := false
	for i, t := range tags {
		if t.Name == tag.Name {
			reduced = append(tags[:i], tags[i+1:]...) // remove from slice
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf(noMatchingTagsWhenRemoveErrorFormat, tag.Name)
	}
	return reduced, nil
}
