package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTag_registerUser(t *testing.T) {
	// given
	tag := Tag{}
	user := User{}

	// when
	tag.registerUser(&user)

	// then
	assert.ElementsMatch(t, []*User{&user}, tag.Users)
}

func TestTag_registerApp(t *testing.T) {
	// given
	tag := Tag{}
	app := App{}

	// when
	tag.registerApp(&app)

	// then
	assert.ElementsMatch(t, []*App{&app}, tag.Apps)
}

func TestTag_broadcast(t *testing.T) {
	t.Run("no users", func(t *testing.T) {
		// given
		tag := Tag{}
		notification := Notification{}

		// when
		err := tag.broadcast(notification)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, NoUserWhenBroadcastErrorMessage, err.Error())
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
		if assert.NoError(t, err) {
			assert.ElementsMatch(t, []*Tag{{Name: "1"}, {Name: "3"}}, tags)
		}
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tags := []*Tag{{Name: "1"}}
		name := "2"

		// when
		tags, err := removeTag(tags, name)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(NoMatchingTagsWhenRemoveErrorFormat, name), err.Error())
		}
	})
}
