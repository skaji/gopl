// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 153.

// Title3 prints the title of an HTML document specified by a URL.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	id := "member_hash_id"
	found := elementByID(doc, id)
	fmt.Println(found)
}

func elementByID(doc *html.Node, id string) *html.Node {
	var found *html.Node
	pre := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					found = n
					return false
				}
			}
		}
		return true
	}
	forEachNode(doc, pre, nil)
	return found
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) bool {
	if pre != nil {
		ok := pre(n)
		if !ok {
			return false
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ok := forEachNode(c, pre, post)
		if !ok {
			return false
		}
	}
	if post != nil {
		ok := post(n)
		if !ok {
			return false
		}
	}
	return true
}
