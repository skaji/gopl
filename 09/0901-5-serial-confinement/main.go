package main

import (
	"fmt"
	"time"
)

// Cake is
type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}

func main() {
	cooked := make(chan *Cake)
	iced := make(chan *Cake)

	// baker1
	go baker(cooked)
	// baker2
	go baker(cooked)

	// icer1
	go icer(iced, cooked)
	// icer1
	go icer(iced, cooked)

	for {
		_ = <-iced
		fmt.Println("got iced cake!")
		time.Sleep(time.Second)
	}
}
