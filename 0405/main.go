package main

import "fmt"

func squish(s []string) []string {
	if len(s) == 0 {
		return s
	}
	j := 0
	for i := 0; i < len(s); i++ {
		if s[j] == s[j+1] {
			copy(s[j:], s[j+1:])
		} else {
			j++
		}
	}
	return s[:j+1]
}

func main() {
	squish([]string{})
	s := []string{"a", "b", "b", "b", "b", "c", "c", "d", "e", "a", "a", "b"}
	s = squish(s)
	fmt.Println(s)
}
