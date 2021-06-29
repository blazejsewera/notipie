package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepository(t *testing.T) {
	// given
	notification := getTestNotification()

	t.Run("save notification", func(t *testing.T) {
		// given
		repo := &MockNotificationRepository{}
		user := &User{repo: repo}

		// when
		user.repo.SaveNotification(notification)

		// then
		assert.ElementsMatch(t, [...]Notification{notification}, repo.Notifications)
	})

	t.Run("get all notifications from repo", func(t *testing.T) {
		// given
		repo := &MockNotificationRepository{Notifications: []Notification{notification}}
		user := &User{repo: repo}

		// when
		notifications := user.repo.GetAllNotifications()

		// then
		assert.ElementsMatch(t, [...]Notification{notification}, notifications)
	})

	t.Run("get all notifications from user", func(t *testing.T) {
		// given
		repo := &MockNotificationRepository{Notifications: []Notification{notification}}
		user := &User{repo: repo}

		// when
		notifications := user.GetAllNotifications()

		// then
		assert.ElementsMatch(t, [...]Notification{notification}, notifications)
	})

	t.Run("get 2 last notifications from user", func(t *testing.T) {
		// given
		notifications := get5TestNotifications()
		repo := &MockNotificationRepository{Notifications: notifications}
		user := &User{repo: repo}

		want := notifications[3:4]

		// when
		have := user.GetLastNotifications(2)

		// then
		assert.ElementsMatch(t, want, have)
	})
}

func TestReceive(t *testing.T) {
	// given
	notification := getTestNotification()

	t.Run("single receive", func(t *testing.T) {
		// given
		repo := &MockNotificationRepository{}
		user := &User{repo: repo}

		// when
		user.Receive(notification)

		// then
		assert.ElementsMatch(t, [...]Notification{notification}, user.GetAllNotifications())
	})

	t.Run("multiple receive - same notification", func(t *testing.T) {
		// given
		repo := &MockNotificationRepository{}
		user := &User{repo: repo}

		// when
		user.Receive(notification)
		user.Receive(notification)

		// then
		assert.ElementsMatch(t, [...]Notification{notification}, user.GetAllNotifications())
	})
}

func TestSubscribeToTag(t *testing.T) {
	// given
	user := &User{}
	tag := &Tag{}

	// when
	user.SubscribeTo(tag)

	// then
	assert.ElementsMatch(t, [...]*Tag{tag}, user.tags)
}
