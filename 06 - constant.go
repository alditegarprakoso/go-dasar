package main

import "fmt"

func main() {
	const firstName string = "Aldi Tegar" // untuk tipe data bisa di deklarasi ataupun tidak
	const lastName string = "Prakoso"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// Constan atau constanta itu nilai nya tidak bisa diubah
	// firstName = "Aldi"

	// Untuk membuat constant /constanta sekaligus
	const (
		religion = "Islam"
		gender   = "Pria"
	)

	fmt.Println(religion)
	fmt.Println(gender)
}
