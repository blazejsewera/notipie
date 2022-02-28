package domain

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/pkg/lib/log"
	"go.uber.org/zap"
	"sync"
)

type Tag struct {
	Name             string
	Users            []*User
	Apps             []*App
	NotificationChan chan Notification
	usersMutex       sync.Mutex
	appsMutex        sync.Mutex
	l                *zap.Logger
}

func NewTag(name string) *Tag {
	return &Tag{Name: name, l: log.For("domain").Named("tag")}
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

	t.l.Debug("started tag", zap.String("name", t.Name))
}

func (t *Tag) broadcast(notification Notification) error {
	if len(t.Users) == 0 {
		t.l.Warn("no users subscribed to tag", zap.String("name", t.Name))
		return fmt.Errorf(NoUserWhenBroadcastErrorMessage)
	}

	for _, user := range t.Users {
		user.NotificationChan <- notification
		t.l.Info(
			"sent notification to user",
			zap.String("username", user.Username),
			zap.String("tagName", t.Name),
			zap.String("notificationID", notification.ID),
		)
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

func (t *Tag) unregisterUser(userID string) {
	var err error

	t.usersMutex.Lock()
	defer t.usersMutex.Unlock()

	t.Users, _, err = removeUser(t.Users, userID)
	if err != nil {
		t.l.Warn("error when unregistering user from tag", zap.Error(err))
		return
	}
	t.l.Info("unregistered app from tag", zap.String("tagName", t.Name), zap.String("userID", userID))
}

func (t *Tag) unregisterApp(appID string) {
	var err error

	t.appsMutex.Lock()
	defer t.appsMutex.Unlock()

	t.Apps, _, err = removeApp(t.Apps, appID)
	if err != nil {
		t.l.Warn("error when unregistering app from tag", zap.Error(err))
		return
	}
	t.l.Info("unregistered app from tag", zap.String("tagName", t.Name), zap.String("appID", appID))
}

const (
	NoUserWhenBroadcastErrorMessage      = "no users to broadcast to"
	NoMatchingTagsWhenRemoveErrorFormat  = "tag %q not found"
	NoMatchingAppsWhenRemoveErrorFormat  = "app %q not found"
	NoMatchingUsersWhenRemoveErrorFormat = "user %q not found"
)

func removeTag(initial []*Tag, name string) (reduced []*Tag, removedTag *Tag, err error) {
	found := false
	reduced = initial
	for i, tag := range reduced {
		if tag.Name == name {
			reduced = append(reduced[:i], reduced[i+1:]...) // remove from slice
			removedTag = tag
			found = true
		}
	}

	if !found {
		err = fmt.Errorf(NoMatchingTagsWhenRemoveErrorFormat, name)
	}
	return
}

func removeApp(initial []*App, id string) (reduced []*App, removedApp *App, err error) {
	found := false
	reduced = initial
	for i, app := range reduced {
		if app.ID == id {
			reduced = append(reduced[:i], reduced[i+1:]...) // remove from slice
			removedApp = app
			found = true
		}
	}

	if !found {
		err = fmt.Errorf(NoMatchingAppsWhenRemoveErrorFormat, id)
	}
	return
}

func removeUser(initial []*User, id string) (reduced []*User, removedUser *User, err error) {
	found := false
	reduced = initial
	for i, user := range reduced {
		if user.ID == id {
			reduced = append(reduced[:i], reduced[i+1:]...) // remove from slice
			removedUser = user
			found = true
		}
	}

	if !found {
		err = fmt.Errorf(NoMatchingUsersWhenRemoveErrorFormat, id)
	}
	return
}
