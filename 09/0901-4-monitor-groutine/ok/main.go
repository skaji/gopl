package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int)
var balances = make(chan int)

// Deposit is
func Deposit(amount int) { deposits <- amount }

// Balance is
func Balance() int { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance = balance + amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller()
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
