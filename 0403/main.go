package main

import "fmt"

// original
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// SIZE is
const SIZE = 10

func reverseArray(s *[SIZE]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	a := [SIZE]int{10, 20, 30}
	reverseArray(&a)
	fmt.Println(a)
}
