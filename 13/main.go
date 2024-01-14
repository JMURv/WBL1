package main

import "fmt"

func main() {
	first, second := 10, 20
	fmt.Printf("Было: %d И %d\n", first, second)

	first, second = second, first
	fmt.Printf("Стало: %d И %d\n", first, second)

}
