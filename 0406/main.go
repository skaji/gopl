package main

import (
	"fmt"
	"unicode"
)

func squish(s []byte) []byte {
	rs := []rune(string(s))
	for i := 0; i < len(rs); i++ {
		if unicode.IsSpace(rs[i]) {
			rs[i] = ' '
		}
	}

	j := 0
	for i := 0; i < len(rs)-1; i++ {
		if rs[j] == ' ' && rs[j+1] == ' ' {
			copy(rs[j:], rs[j+1:])
		} else {
			j++
		}
	}
	return []byte(string(rs[:j]))
}

func main() {
	b := squish([]byte("foo    bar\t\tbaz"))
	fmt.Println(string(b))

}
