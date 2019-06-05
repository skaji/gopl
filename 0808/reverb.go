// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

//!+
func handleConn(c net.Conn) {
	defer c.Close()
	alive := make(chan struct{})
	go func() {
		input := bufio.NewScanner(c)
		for input.Scan() {
			alive <- struct{}{}
			go echo(c, input.Text(), 1*time.Second)
		}
	}()
	for {
		select {
		case <-time.After(10 * time.Second):
			fmt.Println("Forcefully shoutdown")
			return
		case <-alive:
			// ok
		}
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
