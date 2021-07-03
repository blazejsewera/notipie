package domain

import (
	"fmt"
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
		assert.ElementsMatch(t, [...]*User{&user}, tag.Users)
	})

	t.Run("user perspective", func(t *testing.T) {
		// given
		tag := Tag{}
		user := User{}

		// when
		user.SubscribeToTag(&tag)

		// then
		assert.ElementsMatch(t, [...]*User{&user}, tag.Users)
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
	t.Run("app to tag without user", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := getTestApp()
		notification := getTestNotification()

		app.AddTag(&tag)

		// when
		err := app.Send(notification)

		// then
		assert.NotNil(t, err)
		assert.Equal(
			t,
			// TODO: refactor test to use error format constant
			fmt.Sprintf("tags: [ TestTag ] for TestApp#1 did not have registered users when sending %s", notification),
			err.Error(),
		)
	})

	t.Run("app to user - single tag", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := getTestApp()
		user := getTestUser()
		notification := getTestNotification()

		app.AddTag(&tag)
		user.SubscribeToTag(&tag)

		// when
		err := app.Send(notification)

		// then
		assert.Nil(t, err)
		assert.ElementsMatch(t, [...]Notification{notification}, user.GetAllNotifications())
	})
}
