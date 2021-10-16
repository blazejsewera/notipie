package domain_test

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegration_AppToUser(t *testing.T) {
	t.Run("send notification - no tag and no users", func(t *testing.T) {
		// given
		app := NewTestApp()
		notification := NewTestNotification()

		// when
		err := app.Send(notification)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(domain.NoTagsWhenSendErrorFormat, "TestApp", "1", notification), err.Error())
		}
	})

	t.Run("send notification - one tag and no users", func(t *testing.T) {
		// given
		tag := NewTestTag()
		tag.Listen()
		app := NewTestApp()
		notification := NewTestNotification()

		app.AddTag(&tag)

		// when
		err := app.Send(notification)

		// then
		assert.NoError(t, err)
	})

	t.Run("send notification - multiple tags and multiple users", func(t *testing.T) {
		// given
		tag1 := NewTestTag()
		tag1.Listen()
		tag2 := NewTestTag()
		tag2.Listen()

		app := NewTestApp()
		app.AddTag(&tag1)
		app.AddTag(&tag2)

		user1, repo1 := NewTestUserWithAsyncRepo()
		user2, repo2 := NewTestUserWithAsyncRepo()

		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		user1.Listen()
		user2.Listen()

		notification := NewTestNotification()

		// when
		err := app.Send(notification)

		// then
		if assert.NoError(t, err) {

			done1 := repo1.NotificationSaved
			util.AsyncAssert(t, done1).ElementsMatch(
				[]domain.Notification{notification},
				GetAllNotificationsFor(user1),
			)

			done2 := repo2.NotificationSaved
			util.AsyncAssert(t, done2).ElementsMatch(
				[]domain.Notification{notification},
				GetAllNotificationsFor(user2),
			)
		}
	})
}

func TestIntegration_UserToApp(t *testing.T) {
	t.Run("respond with a command after sent notification", func(t *testing.T) {
		// given
		tag := NewTestTag()
		tag.Listen()

		app := NewTestApp()
		app.AddTag(&tag)

		user, _ := NewTestUser()

		user.SubscribeToTag(&tag)
		user.Listen()

		notification := NewTestNotification()

		err := app.Send(notification)

		// when
		if err != nil {
			// user.RespondWithCommand
		}
	})
}
