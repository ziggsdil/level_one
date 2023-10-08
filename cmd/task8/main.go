package main

import (
	"fmt"
	"sync"
)

// todo: проверить действительно ли запись ведется конкуретно
func main() {
	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}

	m := make(map[int]int)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			m[idx] = idx * idx
		}(i)
	}

	wg.Wait()

	fmt.Println(m)
}
