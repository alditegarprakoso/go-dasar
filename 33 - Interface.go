package main

import "fmt"

// Membuat Interface
type HasName interface {
	GetName() string // membuat sebuah inisial function pada interface yang return nya adalah string
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var aldi Person
	aldi.Name = "Aldi Tegar Prakoso"

	SayHello(aldi)

	cat := Animal{Name: "Kucing"}

	SayHello(cat)

}
