package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println("GOMAXPROCS =", runtime.GOMAXPROCS(0))
	for i := 0; i < 1000; i++ {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
