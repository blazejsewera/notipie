package domain_test

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	"github.com/jazzsewera/notipie/core/internal/domain/domain_test/mock"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntegration_AppToUser(t *testing.T) {
	t.Run("send notification - no tag and no users", func(t *testing.T) {
		// given
		app := mock.NewTestApp()
		notification := mock.NewTestNotification()

		// when
		err := app.Send(notification)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(domain.NoTagsWhenSendErrorFormat, "TestApp", "1", notification), err.Error())
		}
	})

	t.Run("send notification - one tag and no users", func(t *testing.T) {
		// given
		tag := mock.NewTestTag()
		tag.Listen()
		app := mock.NewTestApp()
		notification := mock.NewTestNotification()

		app.AddTag(&tag)

		// when
		err := app.Send(notification)

		// then
		assert.NoError(t, err)
	})

	t.Run("send notification - multiple tags and multiple users", func(t *testing.T) {
		// given
		tag1 := mock.NewTestTag()
		tag1.Listen()
		tag2 := mock.NewTestTag()
		tag2.Listen()

		app := mock.NewTestApp()
		app.AddTag(&tag1)
		app.AddTag(&tag2)

		notification := mock.NewTestNotification()

		user1, repo1 := mock.NewTestUserWithAsyncRepo()
		user2, repo2 := mock.NewTestUserWithAsyncRepo()

		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		user1.Listen()
		user2.Listen()

		// when
		err := app.Send(notification)

		// then
		if assert.NoError(t, err) {

			done1 := repo1.NotificationSaved
			util.AsyncAssert(t, done1).ElementsMatch(
				[]domain.Notification{notification},
				mock.GetAllNotificationsFor(user1),
			)

			done2 := repo2.NotificationSaved
			util.AsyncAssert(t, done2).ElementsMatch(
				[]domain.Notification{notification},
				mock.GetAllNotificationsFor(user2),
			)
		}
	})

	t.Run("receive command after sent notification", func(t *testing.T) {
		// TODO: write this test
	})
}
