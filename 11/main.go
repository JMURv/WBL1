package main

import "fmt"

func countOnSlice(data map[string]bool, slice []string) map[string]bool {
	for _, v := range slice {
		data[v] = true
	}
	return data
}

func intersection(slice1, slice2 []string) []string {
	var res []string
	countMap := make(map[string]bool)

	countOnSlice(countMap, slice1)
	countOnSlice(countMap, slice2)

	for k := range countMap {
		res = append(res, k)
	}

	return res
}

func main() {
	slice1 := []string{"first", "second", "first", "third"}
	slice2 := []string{"second", "first", "second"}

	fmt.Println(intersection(slice1, slice2))
}
