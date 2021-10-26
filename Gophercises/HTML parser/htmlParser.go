package htmllink

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {

	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	links := linkNode(doc)

	var attributes []Link
	for _, link := range links {
		attributes = append(attributes, buildLink(link))
	}

	return attributes, nil
}

// linkNode takes in all *html.Node and returns them
func linkNode(n *html.Node) []*html.Node {

	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, linkNode(c)...)
	}

	return ret
}

// buildLink creates the Link struct with href and it's corresponding text field
func buildLink(n *html.Node) Link {

	var ret Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = linkText(n)

	return ret
}

// linkText returns string of all HTML text starting from given root node, stripping extra whitespace
func linkText(n *html.Node) string {

	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += linkText(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}
