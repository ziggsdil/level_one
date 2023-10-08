package main

import "fmt"

// разработать конвейер чисел. даны два канала:
// в первый пишутся числа (x) из массива, во второй
// результат операции x*2, после чего данные из второго
// канала должеы выводиться в stdout

func main() {
	in := make(chan int)
	out := make(chan int)
	nums := []int{1, 2, 3, 5, 6, 7, 8, 9}

	go writer(nums, in)
	go reader(out, in)

	for val := range out {
		fmt.Printf("%d ", val)
	}

}

func writer(nums []int, in chan<- int) {
	for _, num := range nums {
		in <- num
	}
	close(in)
}

func reader(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
