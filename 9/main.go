package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	inputCh := make(chan int)
	outputCh := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)

	// Запускаем первую горутину, которая читает значения из массива и отправляет в канал
	go func() {
		for _, v := range nums {
			inputCh <- v
		}
		close(inputCh)
		wg.Done()
	}()

	// Запускаем вторую горутину, которая читает данные из инпут канала и передает в другой канал умноженными на два
	go func() {
		for num := range inputCh {
			outputCh <- num * 2
		}
		close(outputCh)
		wg.Done()
	}()

	// Запускаем последнюю горутину, которая читает данные из аутпут канала и пишет в stdout
	go func() {
		for result := range outputCh {
			fmt.Println(result)
		}
		wg.Done()
	}()

	wg.Wait()
}
