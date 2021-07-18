package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUserRepository(t *testing.T) {
	// TODO: remove those tests, as they test mocks
	// given
	notification := getTestNotification()

	t.Run("get all notifications from repo", func(t *testing.T) {
		// given
		repo := mockNotificationRepository{Notifications: []Notification{notification}}
		user := User{repo: &repo}

		// when
		notifications := user.GetAllNotifications()

		// then
		assert.ElementsMatch(t, []Notification{notification}, notifications)
	})

	t.Run("get all notifications from user", func(t *testing.T) {
		// given
		repo := mockNotificationRepository{Notifications: []Notification{notification}}
		user := User{repo: &repo}

		// when
		notifications := user.GetAllNotifications()

		// then
		assert.ElementsMatch(t, []Notification{notification}, notifications)
	})

	t.Run("get 2 last notifications from user", func(t *testing.T) {
		// given
		notifications := get5TestNotifications()
		repo := mockNotificationRepository{Notifications: notifications}
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
		repo := mockNotificationRepository{Notifications: notifications}
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
		repo := mockNotificationRepository{Notifications: notifications}
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
		repo := newMockNotificationRepository()
		user.repo = &repo

		// when
		go func() {
			user.Receive(notification)
		}()

		// then
		select {
		case <-repo.NotificationSaved:
			assert.ElementsMatch(t, []Notification{notification}, user.GetAllNotifications())
		case <-time.After(200 * time.Millisecond):
			assert.Fail(t, "notification was not stored in repo after 200ms")
		}
	})

	t.Run("multiple receive - same notification", func(t *testing.T) {
		// given
		// TODO: refactor tests to remove repetition
		user := getTestUser()
		repo := newMockNotificationRepository()
		user.repo = &repo

		// when
		done := make(chan struct{})
		go func() {
			user.Receive(notification)
			user.Receive(notification)
			user.Receive(notification)
			user.Receive(notification)
			done <- struct{}{}
		}()
		<-repo.NotificationSaved

		// then
		select {
		case <-done:
			assert.ElementsMatch(t, []Notification{notification}, user.GetAllNotifications())
		case <-time.After(200 * time.Millisecond):
			assert.Fail(t, "notification was not stored in repo after 200ms")
		}
	})
}

func TestUser_Listen(t *testing.T) {
	// given
	user := getTestUser()
	repo := newMockNotificationRepository()
	user.repo = &repo

	notification := getTestNotification()

	timeout := time.After(200 * time.Millisecond)
	user.Listen()

	// when
	select {
	case user.NotificationChan <- notification:
		// then
		assert.Equal(t, []Notification{notification}, user.GetAllNotifications())
	case <-timeout:
		assert.Fail(t, "user.NotificationChan blocked for over 200ms")
	}
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
		require.NoError(t, err)
		assert.Empty(t, user.tags)
	})

	t.Run("not found", func(t *testing.T) {
		// given
		user := getTestUser()
		tag := getTestTag()
		user.tags = []*Tag{}

		// when
		err := user.UnsubscribeFromTag(tag)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(noMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
	})
}
