package anagram

import (
	"strings"
	"unicode"
)

type ocurrences [26]int

func count(s string) (o ocurrences) {

	s = strings.ToLower(s)

	for _, l := range s {
		if unicode.IsLetter(l) {
			o[l-'a']++
		}
	}

	return
}

func Detect(s string, c []string) (anagrams []string) {

	anagrams = make([]string, 0)

	sc := count(s)

	for _, w := range c {
		if strings.EqualFold(s, w) {
			continue
		}
		cc := count(w)
		if sc == cc {
			anagrams = append(anagrams, w)
		}
	}

	return anagrams
}
