package memo

// Func is
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// Memo is
type Memo struct {
	f     Func
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
	res, ok := m.cache[key]
	if !ok {
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	return res.value, res.err
}
