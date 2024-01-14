package main

import "fmt"

// Human - основная структура
type Human struct {
	FirstName string
	LastName  string
}

// Greeting - метод структуры Human
func (h *Human) Greeting() {
	fmt.Println("Hello, my name is", h.FirstName)
}

// AnonAction - структура в которую анонимно встроена структура Human.
// Анонимность в данном случае означает, что поля и методы структуры-родителя становятся как бы частью структуры Action
type AnonAction struct {
	Human
}

// Walk - метод структуры AnonAction, который может использовать поля и методы из структуры Human
func (a *AnonAction) Walk() {
	fmt.Println(a.FirstName, "is walking")
}

// NamedAction - также структура, но в которую НЕ анонимно встроена структура Human.
type NamedAction struct {
	Person Human
}

// NamedWalk - метод структуры NamedAction, который также может использовать поля и методы из структуры Human
func (a *AnonAction) NamedWalk() {
	fmt.Println(a.FirstName, "is walking")
}

func main() {
	// Используем анонимное встраивание
	anonAction := AnonAction{
		Human: Human{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	// В таком случае не требуется явно указывать родительскую структуру при вызове методов
	anonAction.Greeting()
	anonAction.Walk()

	// Используем НЕ анонимное встраивание
	namedAction := NamedAction{
		Person: Human{
			FirstName: "Fizz",
			LastName:  "Buzz",
		},
	}
	// В таком варианте к полям и методам нужно обращаться явно
	namedAction.Person.Greeting()
}
