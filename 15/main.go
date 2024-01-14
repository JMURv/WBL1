package main

import "strings"

// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// В данном коде есть некоторые проблемы, связанные с неэффективным расходом памяти.
// Операции по конкатенации строк весьма затратны, так как при каждом новом вызове "+=" будет создаваться новая строка
// Поэтому в createHugeString лучше использовать билдер
// Кроме того, в someFunc используется срез от существующей строки. Это не освобождает память от v, а просто делает "окно" в неё
// Поэтому в main'e, после окончания работы программы, я на всякий случай освобождаю justString

var justString string

func createHugeString(size int) string {
	var builder strings.Builder

	for i := 0; i < size; i++ {
		builder.WriteRune('a')
	}

	return builder.String()
}

func someFunc() {
	v := createHugeString(1 << 10) // 1024
	justString = v[:100]
}

func main() {
	someFunc()
	justString = ""
}
