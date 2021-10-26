package main

import (
	"flag"
	"fmt"
	htmlLink "iamargus95/Gophercises/HtmlParser"
	"os"
)

func main() {
	flag.Parse()
	r, _ := os.Open(flag.Arg(0))
	links, err := htmlLink.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
