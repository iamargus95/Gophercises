package main

import (
	"flag"
	"fmt"
	"iamargus95/Gophercises/sitemap/link"
	"log"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://github.com/iamargus95", "Target URL for Site Map")
	flag.Parse()

	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		log.Println("no response body detected")
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)
	for _, link := range links {
		fmt.Println(link)
	}
}
