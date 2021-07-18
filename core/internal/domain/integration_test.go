package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
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
		// TODO: get error from tag
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

		user1, repo1 := newTestUser()
		user2, repo2 := newTestUser()

		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		user1.Listen()
		user2.Listen()

		// when
		err := app.Send(notification)

		// then
		require.NoError(t, err)

		for i := 0; i < 2; i++ {
			select {
			case <-repo1.NotificationSaved:
				assert.ElementsMatch(t, []Notification{notification}, user1.GetAllNotifications())
			case <-repo2.NotificationSaved:
				assert.ElementsMatch(t, []Notification{notification}, user2.GetAllNotifications())
			case <-time.After(200 * time.Millisecond):
				assert.Fail(t, "user1.repo or user2.repo did not save the notification after 200ms")
			}
		}
	})

	t.Run("receive command after sent notification", func(t *testing.T) {
		// TODO: write this test
		t.FailNow()
	})
}
