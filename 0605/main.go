package main

import (
	"bytes"
	"fmt"
)

// BITSIZE is
const BITSIZE = 32 << (^uint(0) >> 63)

// IntSet is
type IntSet struct {
	words []uint
}

// Has is
func (s *IntSet) Has(x int) bool {
	word, bit := x/BITSIZE, uint(x%BITSIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add is
func (s *IntSet) Add(x int) {
	word, bit := x/BITSIZE, uint(x%BITSIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith is
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BITSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", BITSIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len is
func (s *IntSet) Len() int {
	num := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BITSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				num++
			}
		}
	}
	return num
}

// Remove is
func (s *IntSet) Remove(x int) {
	word, bit := x/BITSIZE, uint(x%BITSIZE)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
}

// Clear is
func (s *IntSet) Clear() {
	s.words = nil
}

// Copy is
func (s *IntSet) Copy() *IntSet {
	words := make([]uint, len(s.words))
	copy(words, s.words)
	return &IntSet{words}
}

// AddAll is
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

// IntersectWith is
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

// DifferenceWith is
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			intersect := s.words[i] & tword
			s.words[i] &= ^intersect
		}
	}
}

// SymmetricWith is
func (s *IntSet) SymmetricWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			union := s.words[i] | tword
			intersect := s.words[i] & tword
			s.words[i] = union & ^intersect
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Elems is
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < BITSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, BITSIZE*i+j)
			}
		}
	}
	return elems
}

func main() {
	s := IntSet{}
	s.Add(10)
	s.Add(20)
	s.Add(99)
	s.Add(10)
	for _, e := range s.Elems() {
		fmt.Println(e)
	}
	fmt.Println(s.Elems())
	fmt.Println(s.Len())

	fmt.Println(&s)
	s.Remove(10)
	s.Remove(9)
	fmt.Println(&s)
	s.Clear()
	fmt.Println(&s)
	s.Add(3)
	fmt.Println(&s)
	c := s.Copy()
	fmt.Println(c)
	c.Add(100)
	fmt.Println(&s)
	fmt.Println(c)
	fmt.Println(s.Elems())
}
