package main

import "fmt"

func squish(s []string) []string {
	if len(s) == 0 {
		return s
	}
	j := 0
	for i := 0; i < len(s) && j+1 < len(s); i++ {
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
	s := []string{"a"}
	s = squish(s)
	fmt.Println(s)
}
