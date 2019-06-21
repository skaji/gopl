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

type entry struct {
	result result
	ready  chan struct{}
}

func (e *entry) call(f Func, key string) {
	e.result.value, e.result.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.result
}

func server(requests chan request, f Func) {
	cache := make(map[string]*entry)
	for req := range requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}
