package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	index := 2

	fmt.Println(append(slice[:index], slice[index+1:]...))
}
