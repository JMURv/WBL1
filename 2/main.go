package main

import (
	"fmt"
	"sync"
)

func sqrt(number int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- number * number
}

func main() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int, len(nums))

	// Создаём WaitGroup, чтобы дождаться выполнения всех горутин
	var wg sync.WaitGroup
	wg.Add(len(nums))

	// Запускаем горутину для каждого числа из массива
	for _, v := range nums {
		// Передаём wg по ссылке, чтобы сделать Done строго после вычислений
		// В обратном случае получим deadlock, так как будет попытка записи в закрытый канал
		go sqrt(v, ch, &wg)
	}

	// Закрываем канал, когда все горутины завершат свою работу
	wg.Wait()
	close(ch)

	// Читаем результаты
	for result := range ch {
		fmt.Println(result)
	}
}
