package memo

import "sync"

// Func is
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

// Memo is
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

// New is
func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

// Get is
func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		m.cache[key] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	} else {
		m.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}
