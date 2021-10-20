package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	flag.Parse()

	input := flag.Arg(0)
	delta, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("Enter a valid delta in integer format.")
	}
	output := getCipher(input, delta)
	fmt.Printf("The Cipher of '%s' is %s\n", input, output)

}

func getCipher(input string, delta int) string {
	output := []rune{}
	for _, v := range input {
		if int(v) >= int('A') && int(v) <= int('Z') {
			output = append(output, cipherBig(v, delta))
		} else if int(v) >= int('a') && int(v) <= int('z') {
			output = append(output, cipherSmall(v, delta))
		} else {
			output = append(output, v)
		}
	}
	return string(output)
}

func cipherBig(s rune, delta int) rune {

	ascii := int(s)
	ascii += delta
	if ascii > 90 {
		ascii -= 90
		ascii += 64
	}
	return rune(ascii)
}

func cipherSmall(s rune, delta int) rune {

	ascii := int(s)
	ascii += delta
	if ascii > 122 {
		ascii -= 122
		ascii += 96
	}
	return rune(ascii)
}
