package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// Link is
type Link struct {
	Tag string
	URL string
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Printf("%s: %s\n", link.Tag, link.URL)
	}
}

func visit(links []Link, n *html.Node) []Link {
	if n.Type == html.ElementNode {
		key := ""
		switch n.Data {
		case "a":
			key = "href"
		case "img":
			key = "src"
		case "script":
			key = "src"
		}
		if key != "" {
			for _, a := range n.Attr {
				if a.Key == key {
					links = append(links, Link{Tag: n.Data, URL: a.Val})
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
