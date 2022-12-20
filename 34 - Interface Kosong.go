package main

import "fmt"

// Interface kosong pada function
// Jadi function bisa mengeluarkan return value yang bebas, bisa string, int, boolean, dll
func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups"
	}
}

func main() {
	fmt.Println(Ups(1))
	fmt.Println(Ups(2))
	fmt.Println(Ups(3))

	// Jika ingin diletakan pada sebuah variable, maka type datanya harus interface kosong
	var data interface{} = Ups(4)
	fmt.Println(data)
}
