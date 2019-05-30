// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 221.
//!+

// Netcat1 is a read-only TCP client.
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type result struct {
	time string
	host string
	err  error
}

func main() {

	var wg sync.WaitGroup
	ch := make(chan result)
	for _, port := range os.Args[1:] {
		host := fmt.Sprintf("localhost:%s", port)

		wg.Add(1)
		go func(host string) {
			defer wg.Done()
			conn, err := net.Dial("tcp", host)
			if err != nil {
				ch <- result{host: host, err: err}
				return
			}
			defer conn.Close()
			time, err := bufio.NewReader(conn).ReadString('\n')
			if err != nil {
				ch <- result{host: host, err: err}
				return
			}
			ch <- result{host: host, time: strings.TrimSpace(time)}
		}(host)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		msg := res.time
		if res.err != nil {
			msg = res.err.Error()
		}
		fmt.Printf("%s\t%s\n", res.host, msg)
	}

}

//!-
