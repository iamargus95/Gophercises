package main

import (
	"fmt"
)

func main() {
	in := []string{
		"abcabcbb",
		"bbbbb",
		"pwwkew",
		"",
	}

	for i := range in {
		fmt.Println(longestSubstring(in[i]))
	}
}

func longestSubstring(in string) string {

	var ret string
	inputLength := len(in)
	max := 1
	newUniqueChar := 0
	chars := make([]bool, 122) // z ascii value is 122
	if inputLength == 0 {
		max = 0
	}

	for i := 0; i < inputLength; i++ {
		for j := 0; j < 122; j++ {
			chars[j] = false //initilize array
		}
		count := 0
		chars[in[i]] = true //set ascii idx as true
		count += 1

		for k := i + 1; k < inputLength; k++ {
			if chars[in[k]] {
				break
			}

			chars[in[k]] = true
			count += 1

			if max < count {
				newUniqueChar = i
				max = count
			}
		}
		ret = in[newUniqueChar : newUniqueChar+max]
	}
	return ret
}
