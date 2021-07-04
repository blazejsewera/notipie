package domain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTag_RegisterUser(t *testing.T) {
	t.Run("tag perspective", func(t *testing.T) {
		// given
		tag := Tag{}
		user := User{}

		// when
		tag.RegisterUser(&user)

		// then
		assert.ElementsMatch(t, []*User{&user}, tag.Users)
	})

	t.Run("user perspective", func(t *testing.T) {
		// given
		tag := Tag{}
		user := User{}

		// when
		user.SubscribeToTag(&tag)

		// then
		assert.ElementsMatch(t, []*User{&user}, tag.Users)
	})
}

func TestTag_RegisterApp(t *testing.T) {
	t.Run("tag perspective", func(t *testing.T) {
		// given
		tag := Tag{}
		app := App{}

		// when
		tag.RegisterApp(&app)

		// then
		assert.ElementsMatch(t, []*App{&app}, tag.Apps)
	})

	t.Run("app perspective", func(t *testing.T) {
		// given
		tag := Tag{}
		app := App{}

		// when
		app.AddTag(&tag)

		// then
		assert.ElementsMatch(t, []*App{&app}, tag.Apps)
	})
}

func TestTag_Broadcast(t *testing.T) {
	t.Run("no users", func(t *testing.T) {
		// given
		tag := getTestTag()
		notification := getTestNotification()

		// when
		err := tag.Broadcast(notification)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, noUserWhenBroadcastErrorMessage, err.Error())
		}
	})

	t.Run("multiple users", func(t *testing.T) {
		// given
		tag := getTestTag()
		user1 := getTestUser()
		user2 := getTestUser()
		notification := getTestNotification()

		user1.SubscribeToTag(&tag)
		user2.SubscribeToTag(&tag)

		// when
		err := tag.Broadcast(notification)

		// then
		if assert.NoError(t, err) {
			assert.ElementsMatch(t, []Notification{notification}, user1.GetAllNotifications())
			assert.ElementsMatch(t, []Notification{notification}, user2.GetAllNotifications())
		}
	})
}
