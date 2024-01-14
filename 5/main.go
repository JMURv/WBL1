package main

import (
	"fmt"
	"time"
)

const execTime = 5

func sender(ch chan<- int) {
	defer close(ch) // Закрываем канал после выполнения функции, чтобы не получить панику
	for i := 1; ; i++ {
		ch <- i
	}
}

func main() {
	dataCh := make(chan int)
	timer := time.NewTimer(time.Duration(execTime) * time.Second)

	// Горутина для отправки значений
	go sender(dataCh)

	// Ожидаем значений из канала и завершаем программу по истечении времени
	for {
		select {
		case data, ok := <-dataCh:
			if !ok {
				fmt.Println("Программа завершена")
				return
			}
			fmt.Println("Получено значение:", data)
		case <-timer.C:
			fmt.Printf("Программа завершена по истечении %d секунд\n", execTime)
			return
		}
	}
}
