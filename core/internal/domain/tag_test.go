package domain

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain/mock"
	"github.com/jazzsewera/notipie/core/pkg/lib/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestTag_RegisterUser(t *testing.T) {
	// given
	tag := Tag{}
	user := User{}

	// when
	tag.registerUser(&user)

	// then
	assert.ElementsMatch(t, []*User{&user}, tag.Users)
}

func TestTag_RegisterApp(t *testing.T) {
	// given
	tag := Tag{}
	app := App{}

	// when
	tag.registerApp(&app)

	// then
	assert.ElementsMatch(t, []*App{&app}, tag.Apps)
}

func TestTag_Broadcast(t *testing.T) {
	t.Run("no users", func(t *testing.T) {
		// given
		tag := mock.NewTestTag()
		notification := mock.NewTestNotification()

		// when
		err := tag.broadcast(notification)

		// then
		require.Error(t, err)
		assert.Equal(t, NoUserWhenBroadcastErrorMessage, err.Error())
	})

	t.Run("multiple notifications", func(t *testing.T) {
		// given
		tag := mock.NewTestTag()

		notifications := mock.New5TestNotifications()

		user, repo := mock.NewTestUserWithAsyncRepo()
		user.SubscribeToTag(&tag)
		user.Listen()

		util.AsyncRun(func() {
			for _, notification := range notifications {
				// when
				err := tag.broadcast(notification)

				// then
				require.NoError(t, err)
			}
		})

		done := util.AsyncRun(func() {
			for i := range notifications {
				select {
				// check if all notifications arrived
				case <-repo.NotificationSaved:
					continue
				case <-time.After(200 * time.Millisecond):
					assert.Fail(t, fmt.Sprintf("user.repo did not save the notification no %d after 200ms", i))
				}
			}
		})

		util.AsyncAssert(t, done).ElementsMatch(notifications, user.GetNotifications(0, user.GetNotificationCount()))
	})

	t.Run("multiple users", func(t *testing.T) {
		// given
		tag := mock.NewTestTag()

		notification := mock.NewTestNotification()

		user1, repo1 := mock.NewTestUserWithAsyncRepo()
		user2, repo2 := mock.NewTestUserWithAsyncRepo()

		user1.SubscribeToTag(&tag)
		user2.SubscribeToTag(&tag)

		user1.Listen()
		user2.Listen()

		// when
		err := tag.broadcast(notification)

		// then
		require.NoError(t, err)

		done1 := repo1.NotificationSaved
		util.AsyncAssert(t, done1).ElementsMatch([]Notification{notification}, user1.GetNotifications(0, user1.GetNotificationCount()))
		done2 := repo2.NotificationSaved
		util.AsyncAssert(t, done2).ElementsMatch([]Notification{notification}, user2.GetNotifications(0, user2.GetNotificationCount()))
	})
}

func TestTag_removeTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		tags := []*Tag{
			{Name: "1"},
			{Name: "2"},
			{Name: "3"},
		}
		name := "2"

		// when
		tags, err := removeTag(tags, name)

		// then
		require.NoError(t, err)
		assert.ElementsMatch(t, []*Tag{{Name: "1"}, {Name: "3"}}, tags)
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tags := []*Tag{{Name: "1"}}
		name := "2"

		// when
		tags, err := removeTag(tags, name)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(NoMatchingTagsWhenRemoveErrorFormat, name), err.Error())
	})
}
