package memo

// Func is
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type request struct {
	key      string
	response chan<- result
}

// Memo is
type Memo struct {
	requests chan request
}

// New is
func New(f Func) *Memo {
	requests := make(chan request)
	go server(requests, f)
	return &Memo{requests: requests}
}

// Close is
func (m *Memo) Close() {
	close(m.requests)
}

// Get is
func (m *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	req := request{key: key, response: response}
	m.requests <- req
	res := <-response
	return res.value, res.err
}

func server(requests chan request, f Func) {
	cache := make(map[string]result)
	for req := range requests {
		res, ok := cache[req.key]
		if !ok {
			res.value, res.err = f(req.key)
		}
		req.response <- res
	}
}
