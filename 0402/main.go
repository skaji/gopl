package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	hashName := flag.String("hash", "sha256", "hash func")
	flag.Parse()

	var h hash.Hash
	switch *hashName {
	case "sha384":
		h = sha512.New384()
	case "sha512":
		h = sha512.New()
	default:
		h = sha256.New()
	}

	if _, err := io.Copy(h, os.Stdin); err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", h.Sum(nil))
}
