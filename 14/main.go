package main

import (
	"fmt"
	"reflect"
)

func main() {
	var testVar interface{}
	testVar = 10

	fmt.Println(fmt.Sprintf("%T", testVar))

	// ИЛИ
	switch testVar.(type) {
	case int:
		fmt.Println("Это же int!")
	}

	// ИЛИ
	fmt.Println(reflect.TypeOf(testVar))
}
