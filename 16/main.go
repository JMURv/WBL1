package main

import (
	"fmt"
	"slices"
	"sort"
)

func quicksort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

func main() {
	arr := []int{9, 8, 6, 3, 1, 1}

	slices.Sort(arr)
	// ИЛИ
	quicksort(arr)

	fmt.Println("Отсортированный массив:", arr)
}
