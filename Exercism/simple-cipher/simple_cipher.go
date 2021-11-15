package cipher

import (
	"strings"
	"unicode"
)

type (
	shift    int
	vigenere string
)

func NewCaesar() Cipher {
	return NewShift(3)
}

func NewShift(distance int) Cipher {
	if distance == 0 || distance > 25 || distance < -25 {
		return nil
	}

	return shift(distance)
}

func RotationalCipher(r rune, c shift) rune {
	if unicode.IsLetter(r) {
		return ((unicode.ToLower(r) - 'a' + rune(c) + 26) % 26) + 'a' // add 26 so modulus operation is always positive.
	}
	return -1
}

func (c shift) Encode(input string) string {
	return strings.Map(func(r rune) rune { return RotationalCipher(r, c) }, input)
}

// (shift) Decode decodes an input string shift number of letters to the right.
func (c shift) Decode(input string) string {
	return strings.Map(func(r rune) rune {
		return ((r - 'a' - rune(c) + 26) % 26) + 'a'
	}, input)
}

// NewVigenere returns a Cipher type with a vigenere key to use.
func NewVigenere(key string) Cipher {
	var hasNonA bool
	// The key may only have letters that are lower case, and the key cannot be all 'a'.
	for _, c := range key {
		if !unicode.IsLetter(c) || !unicode.IsLower(c) {
			return nil
		}
		if c != 'a' {
			hasNonA = true
		}
	}
	if !hasNonA || key == "" {
		return nil
	}
	return vigenere(key)
}

// vigParser returns only letters that have been lowercased, for use with vigenere.Encode.
func vigParser(s string) []rune {
	output := []rune{}
	for _, r := range s {
		if unicode.IsLetter(r) {
			output = append(output, unicode.ToLower(r))
		}
	}
	return output
}

// (vigenere) Encode encodes the input string with Vigenere encryption key v.
func (v vigenere) Encode(input string) string {
	var encoded []rune
	parsed := vigParser(input)
	for i, r := range parsed {
		key := rune(v[i%len(v)]) // get relevant key letter from v, 'looping' every len(v) thanks to modulus.
		shift := key - 'a'       // get the letter's place in the alphabet; use the number for shifting.
		shifted := r + shift
		// Avoid using modulus. If the new character is greater than 'z', subtract 26 to 'wrap'.
		if shifted > 'z' {
			shifted -= 26
		}
		encoded = append(encoded, shifted)
		// encoded = append(encoded, ((r-'a'+shift+26)%26)+'a')
	}
	return string(encoded)
}

// vigenere) Decode decodes the input string using Vigenere encryption key v.
func (v vigenere) Decode(input string) string {
	var decoded []rune
	for i, r := range strings.ToLower(input) {
		// See comments in Encode.
		newval := r - (rune(v[i%len(v)]) - 'a')
		if newval < 'a' {
			newval += 26
		}
		decoded = append(decoded, newval)
	}
	return string(decoded)
}
