package atbash

import (
	"strings"
	"unicode"
)

func Atbash(input string) (result string) {
	input = strings.ToLower(input)
	letter := 1
	for _, inputChar := range input {
		if inputChar >= 'a' && inputChar <= 'z' {
			result += string('z' - inputChar + 'a')
		} else {
			if unicode.IsNumber(inputChar) {
				result += string(inputChar)
			} else {
				continue
			}
		}
		if letter%5 == 0 {
			result += " "
		}
		letter++
	}
	return strings.TrimSpace(result)
}
