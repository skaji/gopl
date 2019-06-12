// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.

package main

import (
	"fmt"
	"sync"
)

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance
var withdrow = make(chan int)
var withdrowOK = make(chan bool)

// Deposit is
func Deposit(amount int) { deposits <- amount }

// Balance is
func Balance() int { return <-balances }

// Withdraw is
func Withdraw(amount int) bool {
	withdrow <- amount
	return <-withdrowOK
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-withdrow:
			if amount <= balance {
				balance -= amount
				withdrowOK <- true
			} else {
				withdrowOK <- false
			}
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

func main() {
	Deposit(10)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i, Withdraw(3))
		}(i)
	}
	wg.Wait()
}
