package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	result := camelCase(flag.Arg(0))
	fmt.Printf("The string '%s' has %d words.\n", flag.Arg(0), result)
}

func camelCase(input string) int {
	controlStr := input
	controlStr = strings.ToLower(controlStr)
	count := 1
	for i := range input {
		if controlStr[i] != input[i] {
			count++
		}
	}

	return count
}
