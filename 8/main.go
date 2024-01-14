package main

import (
	"fmt"
)

func setBit(num int64, positionToSet uint, numToSet uint) int64 {
	// Сбрасываем бит в позиции positionToSet.
	mask := ^(int64(1) << positionToSet) // ^(1 << position) инвертирует все биты, чтобы получить маску, где только бит в positionToSet равен 0, а все остальные равны 1.

	// Применяем маску
	num &= mask

	// Устанавливаем бит в новое значение
	num |= int64(numToSet) << positionToSet

	return num
}

func main() {
	var num int64 = 42
	fmt.Printf("Исходное значение: %d\n", num)

	// Устанавливаем 3-й бит в 1
	positionToSet := uint(3)
	numToSet := uint(1)
	num = setBit(num, positionToSet, numToSet)
	fmt.Printf("Установлен %d-й бит в %d: %d\n", positionToSet, numToSet, num)

	// Устанавливаем 5-й бит в 0
	positionToSet = uint(5)
	numToSet = uint(0)
	num = setBit(num, positionToSet, numToSet)
	fmt.Printf("Установлен %d-й бит в %d: %d\n", positionToSet, numToSet, num)
}
