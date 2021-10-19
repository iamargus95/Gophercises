package main

import "fmt"

var (
	alphabet = []rune("abcefghijklmnopqrstuvwxyz") // Same as []rune{'a','b',.....,'z'}
)

func main() {

	cipher := cipher('e', 1, alphabet)
	fmt.Println(string(cipher))
}

func cipher(s rune, delta int, key []rune) rune {
	index := -1
	for i, r := range key {
		if r == s {
			index = i
			break
		}
	}

	if index < 0 {
		panic("index < 0")
	}

	index = (index + delta) % len(key)

	return key[index]
}
