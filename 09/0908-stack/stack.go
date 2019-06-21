package main

import (
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	max := 1000
	if len(os.Args) > 1 {
		var err error
		max, err = strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("launch %d gorutines...\n", max)
	for i := 0; i < max; i++ {
		go func() {
			time.Sleep(1000 * time.Second)
		}()
	}
	time.Sleep(1000 * time.Second)
}
