package domain

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntegration_AppToUser(t *testing.T) {
	t.Run("send notification - no tag and no users", func(t *testing.T) {
		// given
		app := newTestApp()
		notification := newTestNotification()

		// when
		err := app.Send(notification)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(noTagsWhenSendErrorFormat, "TestApp", "1", notification), err.Error())
	})

	t.Run("send notification - one tag and no users", func(t *testing.T) {
		// given
		tag := getTestTag()
		tag.Listen()
		app := newTestApp()
		notification := newTestNotification()

		app.AddTag(&tag)

		// when
		err := app.Send(notification)

		// then
		require.NoError(t, err)
	})

	t.Run("send notification - multiple tags and multiple users", func(t *testing.T) {
		// given
		tag1 := getTestTag()
		tag1.Listen()
		tag2 := getTestTag()
		tag2.Listen()

		app := newTestApp()
		app.AddTag(&tag1)
		app.AddTag(&tag2)

		notification := newTestNotification()

		user1, _ := newTestUserWithAsyncRepo()
		user2, _ := newTestUserWithAsyncRepo()

		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		user1.Listen()
		user2.Listen()

		// when
		err := app.Send(notification)

		// then
		require.NoError(t, err)

		for _, user := range []*User{user1, user2} {
			done := user.repo.(*mockAsyncNotificationRepository).NotificationSaved
			util.AsyncAssert(t, done).ElementsMatch([]Notification{notification}, user.getAllNotifications())
		}
	})

	t.Run("receive command after sent notification", func(t *testing.T) {
		// TODO: write this test
	})
}
