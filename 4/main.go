package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// Определяем кол-во воркеров тут
const wrk = 5

func dataGenerator(dataCh chan int, doneCh chan struct{}) {
	defer close(dataCh)
	for i := 1; ; i++ {
		select {
		case dataCh <- i:
		case <-doneCh:
			return
		}
	}
}

func main() {
	dataCh := make(chan int)            // Канал для передачи данных от главного потока к воркерам
	doneCh := make(chan struct{})       // Канал для сигналов завершения, который передается горутинам
	signalCh := make(chan os.Signal, 1) // Канал ожидания сигнала CTRL+C

	// Запуск воркеров
	var wg sync.WaitGroup
	wg.Add(wrk)
	for i := 1; i <= wrk; i++ {
		go worker(i, dataCh, doneCh, &wg)
	}

	// Горутина для генерации данных и отправки их в канал
	go dataGenerator(dataCh, doneCh)

	// Ожидание сигнала завершения
	signal.Notify(signalCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Ожидание Ctrl+C
	<-signalCh

	// Закрываем канал, отвечающий за закрытие, чтобы сообщить воркерам о завершении работы программы
	close(doneCh)

	// Ожидание завершения воркеров
	wg.Wait()
	fmt.Println("Программа завершена.")
}

// Функция воркера
func worker(id int, dataCh <-chan int, doneCh <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Воркер %d начал работу.\n", id)

	for {
		select {
		case data, _ := <-dataCh:
			fmt.Printf("Воркер %d получил: %d\n", id, data)
		case <-doneCh:
			fmt.Printf("Воркер %d завершает работу по сигналу.\n", id)
			return
		}
	}
}
