package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Client struct {
	baseURL string
	token   string
	HTTP    *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		baseURL: "https://api.github.com",
		token:   token,
		HTTP:    http.DefaultClient,
	}
}

func (c *Client) Do(m string, u string, query map[string]string) ([]byte, error) {
	req, _ := http.NewRequest(m, u, nil)
	req.Header.Add("Authorization", "token "+c.token)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	if len(query) > 0 {
		if m == http.MethodGet {
			q := req.URL.Query()
			for k, v := range query {
				q.Add(k, v)
			}
			req.URL.RawQuery = q.Encode()
		} else {
			req.Header.Add("Content-Type", "application/json")
			b, _ := json.Marshal(query)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
		}
	}
	res, err := c.HTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if !strings.HasPrefix(res.Status, "2") {
		return nil, errors.New(res.Status)
	}
	return body, nil
}

func (c *Client) Create(slug string, title string, body string) (string, error) {
	u := fmt.Sprintf("%s/repos/%s/issues", c.baseURL, slug)
	res, err := c.Do(http.MethodPost, u, map[string]string{"title": title, "body": body})
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (i *Client) Get(slug string, num int) (string, error) {
	// TODO
	return "", nil
}
func (i *Client) Edit(slug string, num int) (string, error) {
	// TODO
	return "", nil
}
func (i *Client) Close(slug string, num int) (string, error) {
	// TODO
	return "", nil
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		fmt.Fprintln(os.Stderr, "miss GITHUB_TOKEN")
		os.Exit(1)
	}
	c := NewClient(token)
	res, err := c.Create("skaji/gopl", "test 1", "this is body")
	fmt.Println(res, err)
}
