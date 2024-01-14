package main

import (
	"fmt"
	"strings"
)

func reverse(input string) string {
	var builder strings.Builder
	runes := []rune(input)

	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}

	return builder.String()
}

func main() {
	inputString := "главрыба"
	reversedString := reverse(inputString)

	fmt.Printf("Исходная строка: %s\n", inputString)
	fmt.Printf("Перевернутая строка: %s\n", reversedString) // абырвалг
}
