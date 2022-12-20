package helper

import "fmt"

var version = 1                    // Variable version tidak bisa digunakan pada package lain karena diawali huruf kecil
var Application = "Belajar Golang" // Variable Application bisa digunakan pada package lain karena diawali huruf besar

// Function SayHello bisa digunakan pada package lain karena diawali dengan huruf besar diawal penamaan function
func SayHello(name string) {
	fmt.Println("Hello", name)
}

// Function sayGoodbye tidak bisa digunakan pada package lain karena diawali dengan huruf kecil diawal penamaan function
func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}
