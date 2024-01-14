package main

import (
	"fmt"
	"time"
)

// Первый способ заключается в использовании отдельного канала для остановки
// В данном случае мы закрываем канал в 33 строке, соответственно в этот момент сигнализируем воркеру, что можно выполнять код в case <-stopCh

func worker(stopCh chan struct{}) {
	for {
		select {
		case <-stopCh:
			fmt.Println("Горутина завершает работу.")
			return
		default:
			fmt.Println("Горутина работает.")
			time.Sleep(time.Second)
		}
	}
}

func main() {
	stopCh := make(chan struct{})

	go worker(stopCh)

	// Даем горутине поработать некоторое время
	time.Sleep(3 * time.Second)

	// Останавливаем горутину
	close(stopCh)
	fmt.Println("Программа завершена.")
}
