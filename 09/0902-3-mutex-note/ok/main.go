package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex
	balance int
)

func deposit(amount int) {
	balance = balance + amount
}

// Deposit is
func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	deposit(amount)
}

// Balance is
func Balance() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

// Withdraw is
func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	// Alice
	go func() {
		defer wg.Done()
		Deposit(200)
		Withdraw(1000)
	}()

	// Bob
	go func() {
		defer wg.Done()
		Deposit(100)
	}()
	wg.Wait()
	fmt.Println(Balance())
}
