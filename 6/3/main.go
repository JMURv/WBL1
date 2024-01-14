package main

import (
	"fmt"
	"sync"
	"time"
)

// Третьим способом встраиваем в структуру Mutex и флаг stop
// Mutex используется для безопасного доступа к флагу, а флаг, собственно, для индикации остановки выполнения горутины

type Worker struct {
	sync.Mutex
	stop bool
}

func (w *Worker) work() {
	for {
		w.Lock()
		if w.stop {
			w.Unlock()
			fmt.Println("Горутина завершает работу.")
			return
		}
		w.Unlock()

		time.Sleep(time.Second)
		fmt.Println("Горутина работает.")
	}
}

func (w *Worker) stopWork() {
	w.Lock()
	defer w.Unlock()
	w.stop = true
}

func main() {
	worker := &Worker{}

	go worker.work()

	time.Sleep(3 * time.Second)

	// Останавливаем горутину
	worker.stopWork()

	fmt.Println("Программа завершена.")
}
