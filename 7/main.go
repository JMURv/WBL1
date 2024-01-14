package main

import (
	"fmt"
	"sync"
)

const writeTries = 100

func worker(index int, data *map[string]int, mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	key := fmt.Sprintf("key-%d", index)

	mu.Lock()
	defer mu.Unlock()

	(*data)[key] = index

	fmt.Printf("Горутина %d записала в мапу значение %d по ключу %s\n", index, index, key)
}

func main() {
	data := make(map[string]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(writeTries)

	// Создание горутин для конкурентной записи в мапу
	for i := 0; i < writeTries; i++ {
		go worker(i, &data, &mu, &wg)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
	fmt.Println("Окончательное состояние мапы:", data)
}
