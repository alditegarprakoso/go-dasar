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

// Function return value, atau function yang mengembalikan nilai
// (name string) adalah sebagai parameter yang harus menggunakan tipe data string
// dan di sebelahnya ada type data string yaitu untuk return value atau pengembalian nilai yang harus dikembalikan dengan type data string
func getHello(name string) string {
	if name == "" {
		return "Maaf, masukkan nama anda dengan benar"
	}
	return "Hello " + name
}

func main() {
	// Function
	sayHello()

	// Function with parameter
	firstName := "Aldi Tegar"
	sayHelloTo(firstName, "Prakoso")

	// Function with parameter and return value
	fmt.Println(getHello(""))

	myName := getHello("Aldi Tegar")
	fmt.Println(myName)

}
