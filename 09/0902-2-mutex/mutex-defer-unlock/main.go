package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

// Deposit is
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance = balance + amount
}

// Balance is
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Alice
	go func() {
		defer wg.Done()
		Deposit(200)
	}()

	// Bob
	go func() {
		defer wg.Done()
		Deposit(100)
	}()
	wg.Wait()
	fmt.Println(Balance())
}
