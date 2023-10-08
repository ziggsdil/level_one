package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"
)

// реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, кототрые читают произвольные
// данные из канала и выводят в stdout. Необходима возможность
// выбоп количества воркеров при старте.

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	workersCount := flag.Int("workers",
		3, "Enter workers count")

	flag.Parse()

	ch := make(chan int)

	for i := 0; i < *workersCount; i++ {
		go writer(ctx, ch)
		go reader(ctx, ch)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println("Программа успешно завершена")
	time.Sleep(1 * time.Second)
	cancel()
}

func writer(ctx context.Context, in chan<- int) {
	someData := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер завершил свою работу")
			return
		default:
			in <- someData
			someData++
			time.Sleep(1 * time.Second)
		}
	}
}

func reader(ctx context.Context, out <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Воркер завершил свою работу.")
			return
		default:
			if val, ok := <-out; ok {
				fmt.Println(val)
			}
		}
	}
}
