package main

import "fmt"

type Cat struct {
	name string
}

func (c Cat) Pet() {
	fmt.Println("miauuu")
}

func (a Cat) Name() string {
	return a.name
}

type Dog struct {
	name string
}

func (d Dog) Pet() {
	fmt.Println("guauuu")
}

func (a Dog) Name() string {
	return a.name
}

type Animal interface {
	Pet()
	Name() string
}

func compliment(a Animal) {
	fmt.Println("God job", a.Name())
	a.Pet()
}

func main() {
	cat := Cat{"felino"}
	compliment(cat)

	dog := Dog{"tachi"}
	compliment(dog)
}
