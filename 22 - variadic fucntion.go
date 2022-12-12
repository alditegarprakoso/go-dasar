package main

import "fmt"

// Variadic function itu adalah dimana sebuah function yang parameternya bisa menampung banyak argument, dan parameter tsb harus ada di posisi paling akhir
// Contoh :
func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	result := sumAll(10, 20, 30, 40, 50)
	fmt.Println(result)

	slice := []int{1, 2, 3, 4, 5}
	hasil := sumAll(slice...) // Jika menggunakan data yang ada pada slice, maka pada argumentnya ditambahkan (...) agar data bisa dimasukkan, kalau di js namanya di spread
	fmt.Println(hasil)
}
