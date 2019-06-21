package memo

import "sync"

// Func is
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Memo is
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

// New is
func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}

// Get is
func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	res, ok := m.cache[key]
	m.mu.Unlock()
	if !ok {
		res.value, res.err = m.f(key)
		m.mu.Lock()
		m.cache[key] = res
		m.mu.Unlock()
	}
	return res.value, res.err
}
