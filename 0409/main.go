package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters

	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		counts[word]++
	}
	if err := in.Err(); err != nil {
		panic(err)
	}
	for k, v := range counts {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
