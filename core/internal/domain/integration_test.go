package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIntegration_AppToUser(t *testing.T) {
	t.Run("send notification - no tag and no users", func(t *testing.T) {
		// given
		app := getTestApp()
		notification := getTestNotification()

		// when
		err := app.Send(notification)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(noTagsWhenSendErrorFormat, "TestApp", "1", notification), err.Error())
		}
	})

	t.Run("send notification - one tag and no users", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := getTestApp()
		notification := getTestNotification()

		app.AddTag(&tag)

		// when
		err := app.Send(notification)

		// then
		if assert.Error(t, err) {
			assert.Equal(
				t,
				fmt.Sprintf(noUsersInTagsWhenSendErrorFormat, []string{"TestTag"}, "TestApp", "1", notification),
				err.Error(),
			)
		}
	})

	t.Run("send notification - multiple tags and multiple users", func(t *testing.T) {
		// given
		tag1 := getTestTag()
		tag2 := getTestTag()

		app := getTestApp()

		notification := getTestNotification()

		user1 := getTestUser()
		repo1 := newMockNotificationRepository()
		user1.repo = &repo1

		user2 := getTestUser()
		repo2 := newMockNotificationRepository()
		user2.repo = &repo2

		timeout := time.After(200 * time.Millisecond)

		app.AddTag(&tag1)
		app.AddTag(&tag2)

		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		user1.Listen()
		user2.Listen()

		// when
		err := app.Send(notification)

		// then
		assert.NoError(t, err)

		for i := 0; i < 2; i++ {
			select {
			case <-repo1.NotificationSaved:
				assert.ElementsMatch(t, []Notification{notification}, user1.GetAllNotifications())
			case <-repo2.NotificationSaved:
				assert.ElementsMatch(t, []Notification{notification}, user2.GetAllNotifications())
			case <-timeout:
				assert.Fail(t, "user1.repo or user2.repo did not save the notification after 200ms")
			}
		}
	})

	t.Run("receive command after sent notification", func(t *testing.T) {
		// TODO: write this test
		t.FailNow()
	})
}
