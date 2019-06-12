// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

package main

import (
	"fmt"
	"sync"
)

// pc[i] is the population count of i.
var pc [256]byte
var once sync.Once

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	once.Do(func() {
		fmt.Println("once called")
		for i := range pc {
			pc[i] = pc[i/2] + byte(i&1)
		}
	})
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i uint64) {
			defer wg.Done()
			fmt.Println(i, PopCount(i))
		}(uint64(i))
	}
	wg.Wait()
}
