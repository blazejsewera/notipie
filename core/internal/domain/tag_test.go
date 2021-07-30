package domain

import (
	"fmt"
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
		tag := getTestTag()
		notification := newTestNotification()

		// when
		err := tag.broadcast(notification)

		// then
		require.Error(t, err)
		assert.Equal(t, noUserWhenBroadcastErrorMessage, err.Error())
	})

	t.Run("multiple notifications", func(t *testing.T) {
		// given
		tag := getTestTag()

		notifications := get5TestNotifications()

		user, repo := newTestUserWithAsyncRepo()
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

		util.AsyncAssert(t, done).ElementsMatch(notifications, user.getAllNotifications())
	})

	t.Run("multiple users", func(t *testing.T) {
		// given
		tag := getTestTag()

		notification := newTestNotification()

		user1, _ := newTestUserWithAsyncRepo()
		user2, _ := newTestUserWithAsyncRepo()

		user1.SubscribeToTag(&tag)
		user2.SubscribeToTag(&tag)

		user1.Listen()
		user2.Listen()

		// when
		err := tag.broadcast(notification)

		// then
		require.NoError(t, err)

		for _, user := range []*User{user1, user2} {
			done := user.repo.(*mockAsyncNotificationRepository).NotificationSaved
			util.AsyncAssert(t, done).ElementsMatch([]Notification{notification}, user.getAllNotifications())
		}
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
		assert.Equal(t, fmt.Sprintf(noMatchingTagsWhenRemoveErrorFormat, name), err.Error())
	})
}
