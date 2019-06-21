package memo

import (
	"fmt"
	"testing"
)

func TestConcurrent(t *testing.T) {
	m := New(func(key string) (interface{}, error) {
		return key, nil
	})

	for i := 0; i < 2; i++ {
		v1, _ := m.Get("v1")
		v2, _ := m.Get("v2")

		fmt.Println("v1", v1)
		fmt.Println("v2", v2)
	}

}
