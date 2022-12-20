package main

import (
	"fmt"
)

// Type Assertion adalah cara untuk mengkonversi interface kosong ke tipe data yang lain
func random() interface{} {
	return "Ups"
}

func main() {
	var result interface{} = random()
	var resultString string = result.(string) // Ini adalah Type Assertion yang dimana untuk mengubah dari interface kosong ke tipe string
	fmt.Println(resultString)

	// Menggunakan switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown Type")
	}

}
