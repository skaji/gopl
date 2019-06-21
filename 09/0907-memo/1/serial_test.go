package memo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestSerial(t *testing.T) {
	httpGetBody := func(url string) (interface{}, error) {
		res, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer res.Body.Close()
		return ioutil.ReadAll(res.Body)
	}
	m := New(httpGetBody)

	urls := []string{
		"https://www.yahoo.co.jp/",
		"https://www.google.co.jp/",
		"https://www.google.co.jp/",
		"https://www.google.com/",
	}

	for _, url := range urls {
		start := time.Now()
		v, err := m.Get(url)
		if err != nil {
			fmt.Println(url, err)
		} else {
			body := v.([]byte) // XXX
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(body))
		}
	}
}
