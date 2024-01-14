package main

import (
	"fmt"
	"time"
)

// Реализовать собственную функцию sleep.

func sleep(d time.Duration) {
	// time.After(duration) создаёт канал, в который по истечению указанного времени отправится текущее время
	// выражение <-time.After(d) блокирует выполнение программы до этого момента
	<-time.After(d)
}

func main() {
	pause := 3
	fmt.Printf("Старт! Пауза: %d секунды\n", pause)
	sleep(time.Second * 3)
	fmt.Println("Пауза завершена.")
}
