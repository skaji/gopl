package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	max := 10
	if len(os.Args) > 1 {
		var err error
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("launch %d gorutines with channel...\n", max)

	first := make(chan int)
	previous := first
	var next chan int
	for i := 0; i < max; i++ {
		next = make(chan int)
		go func(me int, from, to chan int) {
			for v := range from {
				n := v + me
				//fmt.Printf("got %d: %d -> %d\n", me, v, n)
				to <- n
			}
		}(i, previous, next)
		previous = next
	}
	time.Sleep(time.Second)
	first <- 100
	time.Sleep(1000 * time.Second)
}
