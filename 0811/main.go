package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type result struct {
	url    string
	status string
	err    error
}

func request(ctx context.Context, url string) result {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return result{url: url, err: err}
	}
	res, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		return result{url: url, err: err}
	}
	io.Copy(ioutil.Discard, res.Body)
	res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return result{url: url, err: errors.New(res.Status)}
	}
	return result{url: url, status: res.Status}
}

func mirroredQuery() result {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	responses := make(chan result, 3)
	go func() { responses <- request(ctx, "http://www.yahoo.co.jp") }()
	go func() { responses <- request(ctx, "http://www.google.co.jp") }()
	go func() { responses <- request(ctx, "http://www.ikyu.com") }()

	for i := 0; i < 3; i++ {
		if res := <-responses; res.err == nil {
			return res
		}
	}
	return result{url: "N/A", err: errors.New("ALL FAIL")}
}

func main() {
	res := mirroredQuery()
	fmt.Println(res)
}
