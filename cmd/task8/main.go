package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}

	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			mu.Lock()
			m[idx] = idx * idx
			mu.Unlock()
		}(i)
	}

	wg.Wait()

	fmt.Println(m)
}
