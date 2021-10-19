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
		output = append(output, cipher(v, delta))
	}
	return string(output)
}

func cipher(s rune, delta int) rune {

	if int(s) >= int('A') && int(s) <= int('Z') || int(s) >= int('a') && int(s) <= int('z') {
		ascii := int(s)
		ascii += delta
		if ascii > 90 && ascii < 97 {
			ascii -= 90
			ascii += 64
		} else if ascii > 122 {
			ascii -= 122
			ascii += 96
		}
		return rune(ascii)
	}
	return s
}
