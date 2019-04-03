package main

import "fmt"

func squish(s []string) []string {
	j := 0
	for i := 0; i < len(s); i++ {
		if s[j] == s[j+1] {
			copy(s[j:], s[j+1:])
		} else {
			j++
		}
	}
	return s[:j]
}

func main() {
	s := []string{"a", "b", "b", "b", "b", "c", "c", "d", "e", "a"}
	s = squish(s)
	fmt.Println(s)
}
