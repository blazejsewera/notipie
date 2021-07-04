package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository(t *testing.T) {
	// given
	notification := getTestNotification()

	t.Run("save notification", func(t *testing.T) {
		// given
		repo := MockNotificationRepository{}
		user := User{repo: &repo}

		// when
		user.repo.SaveNotification(notification)

		// then
		assert.ElementsMatch(t, []Notification{notification}, repo.Notifications)
	})

	t.Run("get all notifications from repo", func(t *testing.T) {
		// given
		repo := MockNotificationRepository{Notifications: []Notification{notification}}
		user := User{repo: &repo}

		// when
		notifications := user.GetAllNotifications()

		// then
		assert.ElementsMatch(t, []Notification{notification}, notifications)
	})

	t.Run("get all notifications from user", func(t *testing.T) {
		// given
		repo := MockNotificationRepository{Notifications: []Notification{notification}}
		user := User{repo: &repo}

		// when
		notifications := user.GetAllNotifications()

		// then
		assert.ElementsMatch(t, []Notification{notification}, notifications)
	})

	t.Run("get 2 last notifications from user", func(t *testing.T) {
		// given
		notifications := get5TestNotifications()
		repo := MockNotificationRepository{Notifications: notifications}
		user := User{repo: &repo}

		// when
		have := user.GetLastNotifications(2)

		// then
		want := notifications[3:]
		assert.ElementsMatch(t, want, have)
	})

	t.Run("get notifications in specific range", func(t *testing.T) {
		// given
		notifications := get5TestNotifications()
		repo := MockNotificationRepository{Notifications: notifications}
		user := User{repo: &repo}

		// when
		have := user.GetNotifications(1, 3)

		// then
		want := notifications[1:3]
		assert.ElementsMatch(t, want, have)
	})

	t.Run("get notification count", func(t *testing.T) {
		// given
		notifications := get5TestNotifications()
		repo := MockNotificationRepository{Notifications: notifications}
		user := User{repo: &repo}

		// when
		count := user.GetNotificationCount()

		// then
		assert.Equal(t, 5, count)
	})
}

func TestUser_ReceiveNotification(t *testing.T) {
	// given
	notification := getTestNotification()

	t.Run("single receive", func(t *testing.T) {
		// given
		user := getTestUser()

		// when
		user.Receive(notification)

		// then
		assert.ElementsMatch(t, []Notification{notification}, user.GetAllNotifications())
	})

	t.Run("multiple receive - same notification", func(t *testing.T) {
		// given
		user := getTestUser()

		// when
		user.Receive(notification)
		user.Receive(notification)

		// then
		assert.ElementsMatch(t, []Notification{notification}, user.GetAllNotifications())
	})
}

func TestUser_SubscribeToTag(t *testing.T) {
	// given
	user := getTestUser()
	tag := getTestTag()

	// when
	user.SubscribeToTag(&tag)

	// then
	assert.ElementsMatch(t, []*Tag{&tag}, user.tags)
}

func TestUser_UnsubscribeFromTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		user := getTestUser()
		tag := getTestTag()
		user.tags = []*Tag{&tag}

		// when
		err := user.UnsubscribeFromTag(tag)

		// then
		if assert.NoError(t, err) {
			assert.Empty(t, user.tags)
		}
	})
}
