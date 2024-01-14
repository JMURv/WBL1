package main

import (
	"context"
	"fmt"
	"time"
)

// Второй способ остановки горутины заключается в создании контекста WithCancel
// С ним передается коллбэк функция для закрытия канала
// Благодаря <-ctx.Done() можно узнать, что канал завершил работу

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Горутина завершает работу.")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Горутина работает.")
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(3 * time.Second)

	// Останавливаем горутину
	cancel()

	fmt.Println("Программа завершена.")
}
