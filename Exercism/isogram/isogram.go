package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	count := 0
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	word = strings.ReplaceAll(word, "-", "")
	word = strings.ReplaceAll(word, " ", "")

	for _, v := range word {
		count += strings.Count(word, string(v))
	}

	return count <= len(word)
}
