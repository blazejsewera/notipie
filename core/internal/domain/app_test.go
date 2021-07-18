package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestApp_Start(t *testing.T) {
	// given
	app := newTestApp()

	// when
	app.Start()

	// then
	assert.NotNil(t, app.CommandChan)
}

func TestApp_AddTag(t *testing.T) {
	// given
	tag := getTestTag()
	app := newTestApp()

	// when
	app.AddTag(&tag)

	// then
	assert.ElementsMatch(t, []*Tag{&tag}, app.tags)
}

func TestApp_RemoveTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := newTestApp()
		app.tags = []*Tag{&tag}

		// when
		err := app.RemoveTag(tag)

		// then
		require.NoError(t, err)
		assert.Empty(t, app.tags)
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := newTestApp()
		app.tags = []*Tag{{Name: "TestTag2"}}

		// when
		err := app.RemoveTag(tag)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(noMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
	})
}

func TestApp_HandleCommand(t *testing.T) {
	// given
	commandHandler := mockCommandHandler{}
	app := App{commandHandler: &commandHandler}
	command := Command{}
	app.Start()
	timeout := time.After(200 * time.Millisecond)

	// when
	select {
	case app.CommandChan <- command:
		// then
		assert.Equal(t, command, commandHandler.Command)
	case <-timeout:
		assert.Fail(t, "app.CommandChan blocked for over 200ms")
	}
}
