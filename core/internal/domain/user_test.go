package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUser_GetNotifications(t *testing.T) {
	// given
	user, repo := newTestUser()
	notifications := get5TestNotifications()
	repo.Notifications = notifications

	t.Run("get all notifications from user", func(t *testing.T) {
		// when
		have := user.GetAllNotifications()

		// then
		assert.ElementsMatch(t, notifications, have)
	})

	t.Run("get 2 last notifications", func(t *testing.T) {
		// when
		have := user.GetLastNotifications(2)

		// then
		want := notifications[3:]
		assert.ElementsMatch(t, want, have)
	})

	t.Run("get notifications in specific range", func(t *testing.T) {
		// when
		have := user.GetNotifications(1, 3)

		// then
		want := notifications[1:3]
		assert.ElementsMatch(t, want, have)
	})

	t.Run("get notification count", func(t *testing.T) {
		// when
		have := user.GetNotificationCount()

		// then
		want := len(notifications)
		assert.Equal(t, want, have)
	})
}

func TestUser_ReceiveNotification(t *testing.T) {
	// given
	notification := newTestNotification()

	t.Run("single receive", func(t *testing.T) {
		// given
		user, _ := newTestUser()

		// when
		done := asyncRun(func() {
			user.Receive(notification)
		})

		// then
		asyncAssert(t, done).ElementsMatch([]Notification{notification}, user.GetAllNotifications())
	})

	t.Run("multiple receive - same notification", func(t *testing.T) {
		// given
		user, _ := newTestUser()

		// when
		done := asyncRun(func() {
			for i := 0; i < 5; i++ {
				user.Receive(notification)
			}
		})

		// then
		asyncAssert(t, done).ElementsMatch([]Notification{notification}, user.GetAllNotifications())
	})
}

func TestUser_Listen(t *testing.T) {
	// given
	user, _ := newTestUser()

	notification := newTestNotification()

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
	user, _ := newTestUser()
	tag := getTestTag()

	// when
	user.SubscribeToTag(&tag)

	// then
	assert.ElementsMatch(t, []*Tag{&tag}, user.tags)
}

func TestUser_UnsubscribeFromTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		user, _ := newTestUser()
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
		user, _ := newTestUser()
		tag := getTestTag()
		user.tags = []*Tag{}

		// when
		err := user.UnsubscribeFromTag(tag)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(noMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
	})
}

func asyncAssert(t testing.TB, done chan struct{}) *assert.Assertions {
	t.Helper()
	a := assert.New(t)
	select {
	case <-done:
		return a
	case <-time.After(200 * time.Millisecond):
		a.FailNow("test blocked for over 200ms")
		return a
	}
}

func asyncRun(f func()) chan struct{} {
	done := make(chan struct{})
	go func() {
		f()
		done <- struct{}{}
	}()
	return done
}
