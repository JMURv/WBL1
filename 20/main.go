package main

import (
	"fmt"
	"slices"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

const testStr = "snow dog sun"

func main() {
	// Разбиваем строку на слайс по пробелам
	words := strings.Fields(testStr)

	// Переворачиваем слайс
	slices.Reverse(words)

	// Соединяем слайс по разделителю
	fmt.Println(strings.Join(words, " "))
}
