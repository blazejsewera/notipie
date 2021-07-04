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
