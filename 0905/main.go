package main

import (
	"fmt"
	"time"
)

func main() {

	echo := func(name string, from <-chan string, to chan<- string) {
		num := 0
		ticker := time.NewTicker(time.Second)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("%s see %d message/sec\n", name, num)
				num = 0
			case v := <-from:
				num++
				to <- v
			}
		}
	}

	ch1 := make(chan string)
	ch2 := make(chan string)
	go echo("Alice", ch1, ch2)
	go echo("Bob", ch2, ch1)

	ch1 <- "hello"
	time.Sleep(10 * time.Second)
}
