package main

import (
	"flag"
	"fmt"
	"iamargus95/Gophercises/sitemap/link"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "https://github.com/iamargus95", "Target URL for Site Map")
	flag.Parse()

	resp, err := http.Get(*urlFlag)
	if err != nil {
		log.Println("no response body detected")
	}
	defer resp.Body.Close()

	links, _ := link.Parse(resp.Body)

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}

	base := baseURL.String()

	var hrefs []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			hrefs = append(hrefs, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			hrefs = append(hrefs, link.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}
}
