package main

import (
	"fmt"
	"strings"
)

func main() {
	out := expand("$foo $bax fff", func(name string) string {
		return strings.ToUpper(name)
	})
	fmt.Println(out)
}

func expand(s string, f func(string) string) string {
	var out []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '$' {
			var name []byte
			var stop byte
			for {
				i++
				if i >= len(s) {
					break
				}
				next := s[i]
				if next == ' ' || next == '\n' || next == '\t' {
					stop = next
					break
				} else {
					name = append(name, next)
				}
			}
			out = append(out, f(string(name))...)
			out = append(out, stop)
		} else {
			out = append(out, c)
		}
	}
	return string(out)
}
