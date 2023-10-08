package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// реализовать все возможные способы остановки горутины.

// возможные способы остановки: context.WithTimeout
func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx, cancel := context.WithCancel(context.Background())
		wg.Add(1)
		go firstMethod(ctx, wg)
		time.Sleep(2 * time.Second)
		cancel()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		wg.Add(1)
		go secondMethod(wg)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		fmt.Println("Горутина 3 завершена")
		runtime.Goexit()
	}()

	wg.Wait()
	fmt.Println("Main горутина завершена.")
}

// Передача контекста и ожидание его отмены.
func firstMethod(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина 1 завершена.")
			wg.Done()
			return
		}
	}
}

// создание канала и использование его для отправки сигнала завершения.
func secondMethod(wg *sync.WaitGroup) {
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("Горутина 2 завершена")
				wg.Done()
				return
			default:
				fmt.Println("Something")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(3 * time.Second)
	done <- true
}
