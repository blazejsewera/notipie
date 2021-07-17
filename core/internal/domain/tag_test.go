package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
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
		repo1 := newMockNotificationRepository()
		user1.repo = &repo1

		user2 := getTestUser()
		repo2 := newMockNotificationRepository()
		user2.repo = &repo2

		notification := getTestNotification()

		timeout := time.After(200 * time.Millisecond)

		user1.SubscribeToTag(&tag)
		user2.SubscribeToTag(&tag)

		user1.Listen()
		user2.Listen()

		// when
		err := tag.Broadcast(notification)

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
}

func TestTag_removeTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		tags := []*Tag{
			{Name: "1"},
			{Name: "2"},
			{Name: "3"},
		}
		tag := Tag{Name: "2"}
		var err error

		// when
		tags, err = removeTag(tags, tag)

		// then
		if assert.NoError(t, err) {
			assert.ElementsMatch(t, []*Tag{{Name: "1"}, {Name: "3"}}, tags)
		}
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tags := []*Tag{{Name: "1"}}
		tag := Tag{Name: "2"}
		var err error

		// when
		tags, err = removeTag(tags, tag)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(noMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
		}
	})
}
