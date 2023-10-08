package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
)

// программа должна завершаться по нажатию ctrl+c.
// выбрать и обосновать способ завершения работы всех воркеров.

// main: graceful shutdown
func main() {
	// создадим канал для передачи данных о главного потока в воркеры
	dataChannel := make(chan string)

	// канал для завершения всех воркеров
	done := make(chan struct{})

	// wg для синхронизации завершения всех воркеров.
	wg := &sync.WaitGroup{}

	workersCount := 5

	// сразу инициализиурем сколько ожидаем
	wg.Add(workersCount)
	defer wg.Wait() // оиждаем завершение работы всех воркеров.
	// запускаем воркеры
	for i := 1; i <= workersCount; i++ {
		go worker(i, dataChannel, done, wg)
	}

	// обработка сигнала завершения ctrl+c
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		close(done) // закрываем канал завершения воркеров
	}()

	for i := 1; i <= 10; i++ {
		dataChannel <- fmt.Sprintf("some data %d", i)
	}

	// закрываем канал данных после отправил всех данных.
	close(dataChannel)

	<-c // ожидание сигнала завершения.
}

func worker(id int, dataChannel <-chan string, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case data, ok := <-dataChannel:
			if !ok { // канал данных закрыт завершаем работу воркера
				fmt.Printf("Worker %d завершил работу\n", id)
				return
			}
			fmt.Printf("Worker %d получил данные: %s\n", id, data)
		case <-done:
			// получаем сигнал завершения, завершаем работу воркера.
			fmt.Printf("Worker %d завершил работу по сигналу\n", id)
			return
		}
	}
}
