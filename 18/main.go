package main

import (
	"fmt"
	"sync"
)

const workers = 99

type Counter struct {
	sync.Mutex
	value int
}

func (t *Counter) get() int {
	t.Lock()
	defer t.Unlock()
	return t.value
}

func (t *Counter) increment() {
	t.Lock()
	defer t.Unlock()
	t.value += 1
}

func main() {
	c := Counter{}
	var wg sync.WaitGroup

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()
			c.increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Итоговое значение счетчика: %d\n", c.get())
}
