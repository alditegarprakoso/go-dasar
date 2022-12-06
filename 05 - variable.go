package main

import "fmt"

func main() {
	var name string

	name = "Aldi Tegar Prakoso"
	fmt.Println(name)

	name = "NiaSrnt"
	fmt.Println(name)

	// name = true // ketika variable name diisi dengann type data yang tidak sesuai dengan yang dideklarasikan maka akan muncul error

	var address = "Bogor, Indonesia" // Variable bisa di deklarasikan tipe datanya seperti var name atau jika ingin langsung di isi dengan data maka GoLang secara otomatis akan medeteksi tipe data tersebut
	fmt.Println(address)

	phone := "08387654221" // penggunaan "var" untuk mendeklarasi variable sebenarnya tidak wajib, bisa langsung deklarasi menggunakan cara := seperti sintaks disamping
	fmt.Println(phone)

	// Sintaks dibawah berfungsi untuk membuat variable sekaligus
	var (
		firstName = "Aldi Tegar P"
		lastName  = "Nia Sarniati"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
