package main

import "fmt"

func logging() {
	fmt.Println("Function ini dipanggil setelah function sebelumnya selesai")
}

// Function without Defer
func exampleFunc1() {
	fmt.Println("Running funciton exampleFunc1")
	logging()
}

// Function with Defer without Error
func exampleFunc2() {
	defer logging() // Defer digunakan untuk memanggil sebuah function setelah function sebelumnya selesai, walaupun ada error
	fmt.Println("Running funciton exampleFunc2")
}

// Function with Defer with Error
func exampleFunc3(value int) {
	defer logging() // Defer digunakan untuk memanggil sebuah function setelah function sebelumnya selesai, walaupun ada error
	fmt.Println("Running funciton exampleFunc3")
	result := 10 / value
	fmt.Println("Result =", result)
}

func main() {
	exampleFunc1()
	fmt.Println("============================")
	exampleFunc2()
	fmt.Println("============================")
	exampleFunc3(0)
}
