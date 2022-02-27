package domain_test

import (
	"fmt"
	"github.com/blazejsewera/notipie/core/internal/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp_Start(t *testing.T) {
	// given
	app, _ := NewTestApp()

	// when
	app.Start()

	// then
	assert.NotNil(t, app.CommandChan)
}

func TestApp_AddTag(t *testing.T) {
	// given
	tag := NewTestTag()
	app, _ := NewTestApp()

	// when
	app.AddTag(tag)

	// then
	assert.ElementsMatch(t, []*domain.Tag{tag}, app.GetTags())
}

func TestApp_RemoveTag(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		// given
		tag := NewTestTag()
		app, _ := NewTestApp()
		app.AddTag(tag)

		// when
		err := app.RemoveTag(tag.Name)

		// then
		if assert.NoError(t, err) {
			assert.Empty(t, app.GetTags())
		}
	})

	t.Run("not found", func(t *testing.T) {
		// given
		tag := NewTestTag()
		app, _ := NewTestApp()
		app.AddTag(&domain.Tag{Name: "OtherTestTag"})

		// when
		err := app.RemoveTag(tag.Name)

		// then
		if assert.Error(t, err) {
			assert.Equal(t, fmt.Sprintf(domain.NoMatchingTagsWhenRemoveErrorFormat, tag.Name), err.Error())
		}
	})
}
