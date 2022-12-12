package main

import "fmt"

// Deklarasi function
func sayHello() {
	fmt.Println("Hello World")
}

// Function dengan parameter
func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	sayHello()

	firstName := "Aldi Tegar"
	sayHelloTo(firstName, "Prakoso")
}
