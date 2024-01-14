package main

import (
	"fmt"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	a := int64(2 << 20) // 2^21
	b := int64(2 << 21) // 2^22

	res := a * b
	fmt.Printf("Умножение: %d\n", res)

	res = a / b
	fmt.Printf("Деление: %d\n", res)

	res = a + b
	fmt.Printf("Сложение: %d\n", res)

	res = a - b
	fmt.Printf("Вычитание: %d\n", res)
}
