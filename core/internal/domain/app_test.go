package domain

import (
	"fmt"
	"github.com/jazzsewera/notipie/core/internal/domain/mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestApp_Start(t *testing.T) {
	// given
	app := mock.NewTestApp()

	// when
	app.Start()

	// then
	assert.NotNil(t, app.CommandChan)
}

func TestApp_AddTag(t *testing.T) {
	// given
	tag := mock.NewTestTag()
	app := mock.NewTestApp()

	// when
	app.AddTag(&tag)

	// then
	assert.ElementsMatch(t, []*Tag{&tag}, app.tags)
}

func TestApp_RemoveTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		tag := mock.NewTestTag()
		app := mock.NewTestApp()
		app.tags = []*Tag{&tag}

		// when
		err := app.RemoveTag(tag.Name)

		// then
		require.NoError(t, err)
		assert.Empty(t, app.tags)
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tag := mock.NewTestTag()
		app := mock.NewTestApp()
		app.tags = []*Tag{{Name: "TestTag2"}}

		// when
		err := app.RemoveTag(tag.Name)

		// then
		require.Error(t, err)
		assert.Equal(t, fmt.Sprintf(NoMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
	})
}

func TestApp_HandleCommand(t *testing.T) {
	// given
	commandHandler := mock.CommandHandler{}
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
