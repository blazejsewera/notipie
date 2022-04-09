package fp_test

import (
	"github.com/blazejsewera/notipie/core/pkg/lib/fp"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	// given
	input := []int{1, 2, 3}
	add1 := func(i int) int { return i + 1 }
	expected := []int{2, 3, 4}

	// when
	actual := fp.Map(add1, input)

	// then
	assert.ElementsMatch(t, expected, actual)
}
