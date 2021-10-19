package logs

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	isPresent := strings.IndexAny(log, "❗🔍☀")
	if isPresent < 0 {
		return "default"
	}
	return map[rune]string{'☀': "weather", '❗': "recommendation", '🔍': "search"}[[]rune(log)[isPresent]]
}

// Replace replaces all occurances of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	return strings.ReplaceAll(log, fmt.Sprintf("%c", oldRune), fmt.Sprintf("%c", newRune))
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
