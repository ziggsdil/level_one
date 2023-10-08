package main

import "fmt"

type Set struct {
	m map[string]bool
}

func NewSet(l int) *Set {
	return &Set{
		m: make(map[string]bool, l),
	}
}

func (s *Set) Delete(key string) {
	delete(s.m, key)
}

func (s *Set) Add(key string, val bool) {
	s.m[key] = val
}

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}
	set := NewSet(len(sequence))

	for _, s := range sequence {
		set.Add(s, true)
	}
	fmt.Println(set)
	set.Delete("cat")
	fmt.Println(set)
}
