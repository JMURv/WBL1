package main

import "fmt"

func createSet(strings []string) []string {
	var set []string
	countMap := make(map[string]bool)

	for _, str := range strings {
		countMap[str] = true
	}

	// Обратно собираем из мапы слайс
	for k := range countMap {
		set = append(set, k)
	}

	return set
}

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(createSet(slice))
}
