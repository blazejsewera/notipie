package util

import "strings"

// SplitLines - method useful when iterating over lines in a string.
// It returns a slice of strings, each one being a single line.
// If the last couple of lines are newline characters, they are trimmed.
func SplitLines(s string) []string {
	return strings.Split(strings.TrimRight(s, "\n"), "\n")
}
