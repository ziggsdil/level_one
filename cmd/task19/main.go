package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.*/

type Counter struct {
	Val int64
}

func main() {
	c := &Counter{}
	wg := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go c.AddOne(int64(1), wg)
	}
	wg.Wait()
	fmt.Println(c.Val)
}

// используя пакет atomic, для атомарного доступа к переменной
func (c *Counter) AddOne(num int64, wg *sync.WaitGroup) {
	atomic.AddInt64(&c.Val, num)
	wg.Done()
}
