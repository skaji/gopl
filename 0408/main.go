package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func name(c rune) string {
	if unicode.IsControl(c) {
		return "control"
	}
	if unicode.IsDigit(c) {
		return "digit"
	}
	if unicode.IsGraphic(c) {
		return "graphic"
	}
	if unicode.IsLetter(c) {
		return "letter"
	}
	if unicode.IsLower(c) {
		return "lower"
	}
	if unicode.IsMark(c) {
		return "mark"
	}
	if unicode.IsNumber(c) {
		return "number"
	}
	if unicode.IsPrint(c) {
		return "printable"
	}
	if !unicode.IsPrint(c) {
		return "not"
	}
	if unicode.IsPunct(c) {
		return "punct"
	}
	if unicode.IsSpace(c) {
		return "space"
	}
	if unicode.IsSymbol(c) {
		return "symbol"
	}
	if unicode.IsTitle(c) {
		return "title"
	}
	if unicode.IsUpper(c) {
		return "upper"
	}
	panic("oops")
}

func main() {
	counts := make(map[string]int) // counts of Unicode characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		counts[name(r)]++
	}
	for k, v := range counts {
		fmt.Printf("%s -> %d\n", k, v)
	}
}
