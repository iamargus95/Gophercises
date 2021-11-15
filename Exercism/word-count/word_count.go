package wordcount

import "strings"

// Frequency defines a type map with a string key and an int value.
type Frequency map[string]int

// WordCount counts the occurences of each word in a string.
func WordCount(s string) Frequency {

	s = strings.ToLower(s)
	replacer := strings.NewReplacer(".", " ", ",", " ", ":", " ", ";", " ", "!", " ", "?", " ", "&", " ", "@", " ", "$", " ", "%", " ", "^&", " ")

	s = replacer.Replace(s)

	wordMap := Frequency{}
	for _, v := range strings.Fields(s) {
		v = strings.Trim(v, "'")
		wordMap[v]++
	}

	return wordMap
}
