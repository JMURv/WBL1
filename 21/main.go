package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) GetName() string {
	return p.Name
}

// TargetInterface является тем интерфейсом, который должен мэтчится с Person
type TargetInterface interface {
	GetText() string
}

// Создаем адаптер, который реализует необходимый метод GetText
// В адаптер встраивается структура Person. Таким образом мы получаем возможность вызывать метод интерфейса, не адаптируя изначальную структуру Person.
type Adapter struct {
	Person *Person
}

func (a *Adapter) GetText() string {
	return a.Person.GetName()
}

func main() {
	person := &Person{Name: "John"}
	adapter := &Adapter{Person: person}

	// Используем адаптер для вызова метода интерфейса
	fmt.Println(adapter.GetText())
}
