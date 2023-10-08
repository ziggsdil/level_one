package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	//firstSolution(nums)
	secondSolution(nums)
}

// firstSolution using WG
func firstSolution(nums []int) {
	wg := &sync.WaitGroup{}
	mu := sync.Mutex{}

	sum := 0
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			sum += num * num
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(sum)
}

// secondSolution using atomic
func secondSolution(nums []int) {
	wg := &sync.WaitGroup{}
	var sumAt atomic.Int64
	for _, num := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sumAt.Add(int64(num * num))
		}(num)
	}

	wg.Wait()
	fmt.Println(sumAt.Load())
}
