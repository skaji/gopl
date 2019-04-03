package main

import (
	"crypto/sha256"
	"fmt"
)

// 4.1
func diffNum(a, b string) int {
	aSha := sha256.Sum256([]byte(a))
	bSha := sha256.Sum256([]byte(b))

	num := 0
	for i := 0; i < sha256.Size; i++ {
		ba := aSha[i]
		bb := bSha[i]
		if ba != bb {
			num++
		}
	}
	return num
}

func main() {
	fmt.Println(diffNum("a", "a"))
}
