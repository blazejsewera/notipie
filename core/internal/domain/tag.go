package domain

import (
	"fmt"
	"sync"
)

type Tag struct {
	Name             string
	Users            []*User
	Apps             []*App
	NotificationChan chan Notification
	usersMutex       sync.Mutex
	appsMutex        sync.Mutex
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
		return fmt.Errorf(NoUserWhenBroadcastErrorMessage)
	}

	for _, user := range t.Users {
		user.NotificationChan <- notification
	}
	return nil
}

func (t *Tag) registerUser(user *User) {
	t.usersMutex.Lock()
	defer t.usersMutex.Unlock()
	t.Users = append(t.Users, user)
}

func (t *Tag) registerApp(app *App) {
	t.appsMutex.Lock()
	defer t.appsMutex.Unlock()
	t.Apps = append(t.Apps, app)
}

const (
	NoUserWhenBroadcastErrorMessage     = "no users to broadcast to"
	NoMatchingTagsWhenRemoveErrorFormat = "tag %q not found"
)

func removeTag(tags []*Tag, name string) ([]*Tag, error) {
	var reduced []*Tag
	found := false
	for i, tag := range tags {
		if tag.Name == name {
			reduced = append(tags[:i], tags[i+1:]...) // remove from slice
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf(NoMatchingTagsWhenRemoveErrorFormat, name)
	}
	return reduced, nil
}
