package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitLines(t *testing.T) {
	str1 := `First
Second
Third`
	assert.Equal(t, []string{"First", "Second", "Third"}, SplitLines(str1))

	str2 := "First"
	assert.Equal(t, []string{"First"}, SplitLines(str2))

	str3 := `First

Second

Third`
	assert.Equal(t, []string{"First", "", "Second", "", "Third"}, SplitLines(str3))

	str4 := ""
	assert.Equal(t, []string{""}, SplitLines(str4))

	str5 := `First

`
	assert.Equal(t, []string{"First"}, SplitLines(str5))
}
