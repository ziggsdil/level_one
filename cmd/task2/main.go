package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}

	//firstSolution(arr)
	//secondSolution(arr)
	thirdSolution(arr)
}

// firstSolution using waitgroups
func firstSolution(nums [5]int) {
	wg := &sync.WaitGroup{}
	res := make([]int, len(nums))
	for idx, num := range nums {
		wg.Add(1)
		go func(idx, num int) {
			defer wg.Done()
			res[idx] = num * num
		}(idx, num)
	}
	wg.Wait()
	fmt.Println(res)
}

// secondSolution: using buffer chanel
func secondSolution(nums [5]int) {
	res := make(chan int, len(nums))
	sqNums := make([]int, len(nums))

	for _, num := range nums {
		go func(num int) {
			sq := num * num
			res <- sq
		}(num)
	}

	defer close(res)

	for i := 0; i < len(nums); i++ {
		sq := <-res
		sqNums[i] = sq
	}

	fmt.Println(sqNums)
}

func thirdSolution(nums [5]int) {
	sqNums := make([]int, len(nums))
	ch := make(chan int)
	defer close(ch)

	for _, num := range nums {
		go func(num int) {
			sq := num * num
			ch <- sq
		}(num)
	}

	idx := 0
	go func() {
		for num := range ch {
			sqNums[idx] = num
			idx++
		}
	}()

	time.Sleep(1 * time.Second)
	fmt.Println(sqNums)
}
