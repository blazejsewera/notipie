package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	FieldS string
	FieldI int
	FieldB bool
	NestedStruct
}

type NestedStruct struct {
	Field string
}

func TestMerge(t *testing.T) {
	// given
	base := testStruct{
		FieldS:       "s",
		FieldI:       -1,
		FieldB:       true,
		NestedStruct: NestedStruct{Field: "nested old"},
	}
	patch := testStruct{
		FieldI:       1,
		NestedStruct: NestedStruct{Field: "nested new"},
	}
	expected := testStruct{
		FieldS:       "s",
		FieldI:       1,
		FieldB:       true,
		NestedStruct: NestedStruct{Field: "nested new"},
	}

	// when
	actual := Merge(base, patch)

	// then
	assert.Equal(t, expected, actual)
}
