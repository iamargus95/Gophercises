package encode

import (
	"strconv"
	"strings"
	"unicode"
)

func RunLengthEncode(s string) string {
	var encoded string
	for len(s) > 0 {
		letter := s[0]
		slen := len(s)
		s = strings.TrimLeft(s, string(letter))
		if n := slen - len(s); n > 1 {
			encoded += strconv.Itoa(n)
		}
		encoded += string(letter)
	}
	return encoded
}

func RunLengthDecode(s string) string {
	var decoded string
	for len(s) > 0 {
		i := strings.IndexFunc(s, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		n := 1
		if i != 0 {
			n, _ = strconv.Atoi(s[:i])
		}
		decoded += strings.Repeat(string(s[i]), n)
		s = s[i+1:]
	}
	return decoded
}
