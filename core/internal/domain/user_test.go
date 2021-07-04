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
		notifications := user.repo.GetAllNotifications()

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
		want := notifications[3:4]
		assert.ElementsMatch(t, want, have)
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
