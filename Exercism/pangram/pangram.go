package pangram

import "strings"

func IsPangram(input string) bool {

	alphabet := "abcdefghijklmnopqrstuvwxyz"

	input = strings.ToLower(input)
	input = strings.ReplaceAll(input, " ", "")

	for _, v := range alphabet {
		if !strings.ContainsRune(input, v) {
			return false
		}
	}

	return true
}
