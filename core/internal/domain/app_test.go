package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_Send(t *testing.T) {
	t.Run("no tag and no users", func(t *testing.T) {
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

	t.Run("one tag and no users", func(t *testing.T) {
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

	t.Run("multiple tags and multiple users", func(t *testing.T) {
		// given
		tag1 := getTestTag()
		tag2 := getTestTag()
		app := getTestApp()
		notification := getTestNotification()
		user1 := getTestUser()
		user2 := getTestUser()

		app.AddTag(&tag1)
		app.AddTag(&tag2)
		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		// when
		err := app.Send(notification)

		// then
		if assert.NoError(t, err) {
			assert.ElementsMatch(t, []Notification{notification}, user1.GetAllNotifications())
			assert.ElementsMatch(t, []Notification{notification}, user2.GetAllNotifications())
		}
	})
}
