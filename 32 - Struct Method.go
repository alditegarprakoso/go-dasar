package main

import "fmt"

type Student struct {
	Name, Address string
	Age           int
}

// Membuat struct method / function
func (student Student) sayHi(name string) {
	fmt.Println("Hello", name, ", my name is", student.Name)
}

func main() {
	var aldi Student
	aldi.Name = "Aldi Tegar Prakoso"

	aldi.sayHi("Nia")
}
