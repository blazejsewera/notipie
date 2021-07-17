package domain

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestApp_Start(t *testing.T) {
	// given
	app := getTestApp()

	// when
	app.Start()

	// then
	assert.NotNil(t, app.commandChan)
}

func TestApp_AddTag(t *testing.T) {
	// given
	tag := getTestTag()
	app := getTestApp()

	// when
	app.AddTag(&tag)

	// then
	assert.ElementsMatch(t, []*Tag{&tag}, app.tags)
}

func TestApp_RemoveTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := getTestApp()
		app.tags = []*Tag{&tag}

		// when
		err := app.RemoveTag(tag)

		// then
		if assert.NoError(t, err) {
			assert.Empty(t, app.tags)
		}
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tag := getTestTag()
		app := getTestApp()
		app.tags = []*Tag{{Name: "TestTag2"}}

		// when
		err := app.RemoveTag(tag)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(noMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
		}
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
	case app.commandChan <- command:
		// then
		assert.Equal(t, command, commandHandler.Command)
	case <-timeout:
		assert.Fail(t, "app.commandChan blocked for over 200ms")
	}
}
