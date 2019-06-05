// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"gopl.io/ch5/links"
)

type link struct {
	url   string
	depth int
}

func newLinks(urls []string, depth int) []link {
	out := make([]link, 0, len(urls))
	for _, url := range urls {
		out = append(out, link{url: url, depth: depth})
	}
	return out
}

func crawl(l link) []link {
	list, err := links.Extract(l.url)
	if err != nil {
		log.Print(err)
		return nil
	}
	out := []string{}
	for _, l := range list {
		if strings.HasPrefix(l, "http") {
			out = append(out, l)
		}
	}
	fmt.Printf("CRAWL depth %d, url %s, links %d\n", l.depth, l.url, len(out))
	return newLinks(out, l.depth+1)
}

//!+
func main() {
	maxDepth := flag.Int("depth", 1, "max depth")
	flag.Parse()

	worklist := make(chan []link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan link) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- newLinks(flag.Args(), 0) }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 1; i++ {
		go func() {
			for l := range unseenLinks {
				if l.depth > *maxDepth {
					fmt.Printf("SKIP  depth %d, url %s\n", l.depth, l.url)
					continue
				}
				foundLinks := crawl(l)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, l := range list {
			if !seen[l.url] {
				seen[l.url] = true
				unseenLinks <- l
			}
		}
	}
}

//!-
