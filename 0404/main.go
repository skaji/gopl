package main

import "fmt"

func rotate(s []string) []string {
	if len(s) == 0 {
		return s
	}
	return append(s[1:], s[0])
}

func main() {
	s := []string{"a", "b", "c"}
	s = rotate(s)
	fmt.Println(s)
}
