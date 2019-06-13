package main

import (
	"fmt"
	"sync"
)

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func init() {
	sema <- struct{}{}
}

// Deposit is
func Deposit(amount int) {
	<-sema
	balance = balance + amount
	sema <- struct{}{}
}

// Balance is
func Balance() int {
	<-sema
	b := balance
	sema <- struct{}{}
	return b
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
