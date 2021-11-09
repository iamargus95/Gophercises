package main

import (
	"flag"
	"fmt"
	"iamargus95/Gophercises/sitemap/link"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	urlFlag := flag.String("url", "https://github.com/iamargus95", "Target URL for Site Map")
	flag.Parse()

	pages := sitemap(*urlFlag)
	for _, page := range pages {
		fmt.Println(page)
	}
}

func withPrefix(prefix string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, prefix)
	}
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func hrefBuilder(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	var hrefs []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			hrefs = append(hrefs, base+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			hrefs = append(hrefs, link.Href)
		}
	}

	return hrefs
}

func sitemap(site string) []string {
	resp, err := http.Get(site)
	if err != nil {
		log.Println("no response body detected")
	}

	reqURL := resp.Request.URL
	baseURL := &url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()

	return filter(hrefBuilder(resp.Body, base), withPrefix(base))
}
