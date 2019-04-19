package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, tag := range []string{"p", "div", "span"} {
		fmt.Printf("%s %d\n", tag, findNum(tag, doc))
	}
}

func findNum(tag string, n *html.Node) int {
	num := 0
	if n.Type == html.ElementNode && n.Data == tag {
		num++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		num += findNum(tag, c)
	}
	return num
}
