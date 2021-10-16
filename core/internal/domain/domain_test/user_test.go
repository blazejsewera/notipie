package domain_test

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestUser_GetNotifications(t *testing.T) {
	// given
	user, repo := NewTestUser()
	notifications := New5TestNotifications()
	repo.Notifications = notifications

	t.Run("get all notifications from user", func(t *testing.T) {
		// when
		have := user.GetNotifications(0, user.GetNotificationCount())

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
	app, _ := NewTestApp()
	notification := NewTestNotification(app)

	t.Run("single receive", func(t *testing.T) {
		// given
		user, _ := NewTestUser()

		// when
		done := util.AsyncRun(func() {
			user.Receive(notification)
		})

		// then
		util.AsyncAssert(t, done).ElementsMatch([]domain.Notification{notification}, user.GetNotifications(0, user.GetNotificationCount()))
	})

	t.Run("multiple receive - same notification", func(t *testing.T) {
		// given
		user, _ := NewTestUser()

		// when
		done := util.AsyncRun(func() {
			for i := 0; i < 5; i++ {
				user.Receive(notification)
			}
		})

		// then
		util.AsyncAssert(t, done).ElementsMatch([]domain.Notification{notification}, user.GetNotifications(0, user.GetNotificationCount()))
	})
}

func TestUser_Listen(t *testing.T) {
	// given
	user, _ := NewTestUser()

	app, _ := NewTestApp()
	notification := NewTestNotification(app)

	timeout := time.After(200 * time.Millisecond)
	user.Listen()

	// when
	select {
	case user.NotificationChan <- notification:
		// then
		assert.Equal(t, []domain.Notification{notification}, user.GetNotifications(0, user.GetNotificationCount()))
	case <-timeout:
		assert.Fail(t, "user.NotificationChan blocked for over 200ms")
	}
}

func TestUser_SubscribeToTag(t *testing.T) {
	// given
	user, _ := NewTestUser()
	tag := NewTestTag()

	// when
	user.SubscribeToTag(&tag)

	// then
	assert.ElementsMatch(t, []*domain.Tag{&tag}, user.Tags)
}

func TestUser_UnsubscribeFromTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		user, _ := NewTestUser()
		tag := NewTestTag()
		user.Tags = []*domain.Tag{&tag}

		// when
		err := user.UnsubscribeFromTag(tag.Name)

		// then
		require.NoError(t, err)
		assert.Empty(t, user.Tags)
	})

	t.Run("not found", func(t *testing.T) {
		// given
		user, _ := NewTestUser()
		tag := NewTestTag()
		user.Tags = []*domain.Tag{}

		// when
		err := user.UnsubscribeFromTag(tag.Name)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(domain.NoMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
	})
}