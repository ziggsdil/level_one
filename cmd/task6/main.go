package main

import (
	"context"
	"flag"
	"fmt"
	"time"
)

// разработать программу, которая будет последовательно отправилять
// значения в канал, а с другой стороны канала - читать. По
// истечению N секунд программа должна завершаться.

func main() {
	timeout := flag.Int("timeout", 5, "Enter timeout")
	flag.Parse()

	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(*timeout)*time.Second,
	)
	defer cancel()

	in := make(chan int)

	go writer(ctx, in)
	go reader(in)

	select {
	case <-ctx.Done():
		fmt.Printf("Завершение программы по timeout: %d\n", *timeout)
		return
	}
}

func writer(ctx context.Context, in chan<- int) {
	val := 123
	//defer close(in)
	for {
		select {
		case <-ctx.Done(): // для безопасного завершения работы функциии
			return
		default:
			in <- val
			val += 123
		}
	}
}

func reader(out <-chan int) {
	for val := range out {
		fmt.Println(val)
	}
}
