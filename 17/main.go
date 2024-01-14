package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	target := 3

	// sort.Search() использует бинарный поиск под капотом
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

	if index < len(arr) {
		fmt.Printf("Элемент %d найден в позиции %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
