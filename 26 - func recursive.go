package main

import "fmt"

// Function yang memanggil dirinya sendiri
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	recursive := factorialRecursive(5)

	fmt.Println(recursive)
}
