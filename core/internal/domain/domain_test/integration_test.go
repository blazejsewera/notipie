package domain_test

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain"
	mock2 "github.com/jazzsewera/notipie/core/internal/domain/mock"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIntegration_AppToUser(t *testing.T) {
	t.Run("send notification - no tag and no users", func(t *testing.T) {
		// given
		app := mock2.NewTestApp()
		notification := mock2.NewTestNotification()

		// when
		err := app.Send(notification)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(domain.NoTagsWhenSendErrorFormat, "TestApp", "1", notification), err.Error())
	})

	t.Run("send notification - one tag and no users", func(t *testing.T) {
		// given
		tag := mock2.NewTestTag()
		tag.Listen()
		app := mock2.NewTestApp()
		notification := mock2.NewTestNotification()

		app.AddTag(&tag)

		// when
		err := app.Send(notification)

		// then
		require.NoError(t, err)
	})

	t.Run("send notification - multiple tags and multiple users", func(t *testing.T) {
		// given
		tag1 := mock2.NewTestTag()
		tag1.Listen()
		tag2 := mock2.NewTestTag()
		tag2.Listen()

		app := mock2.NewTestApp()
		app.AddTag(&tag1)
		app.AddTag(&tag2)

		notification := mock2.NewTestNotification()

		user1, repo1 := mock2.NewTestUserWithAsyncRepo()
		user2, repo2 := mock2.NewTestUserWithAsyncRepo()

		user1.SubscribeToTag(&tag1)
		user1.SubscribeToTag(&tag2)
		user2.SubscribeToTag(&tag1)

		user1.Listen()
		user2.Listen()

		// when
		err := app.Send(notification)

		// then
		require.NoError(t, err)

		done1 := repo1.NotificationSaved
		util.AsyncAssert(t, done1).ElementsMatch([]domain.Notification{notification}, getAllNotificationsFor(user1))

		done2 := repo2.NotificationSaved
		util.AsyncAssert(t, done2).ElementsMatch([]domain.Notification{notification}, getAllNotificationsFor(user2))
	})

	t.Run("receive command after sent notification", func(t *testing.T) {
		// TODO: write this test
	})
}

func getAllNotificationsFor(u *domain.User) []domain.Notification {
	return u.GetNotifications(0, u.GetNotificationCount())
}
