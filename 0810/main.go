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
	"context"
	"fmt"
	"os"
)

func crawl(ctx context.Context, url string) []string {
	fmt.Println(url)
	list, err := Extract(ctx, url)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return list
}

//!+
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		os.Stdin.Read(make([]byte, 1))
		cancel()
	}()
	run(ctx)
}

func run(ctx context.Context) {

	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(ctx, link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for {
		select {
		case <-ctx.Done():
			close(unseenLinks)
			return
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}

	}
}

//!-
