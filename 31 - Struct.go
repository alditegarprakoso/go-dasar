package main

import "fmt"

// Membuat Struct

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	// Menambahkan data struct

	var aldi Customer
	aldi.Name = "Aldi Tegar Prakoso"
	aldi.Address = "Bogor, Indonesia"
	aldi.Age = 23

	fmt.Println("=============================================")
	fmt.Println(aldi)
	fmt.Println("=============================================")
	fmt.Println(aldi.Name)
	fmt.Println(aldi.Address)
	fmt.Println(aldi.Age)
	fmt.Println("=============================================")
	// Menggunakan Cara Struct Literals
	// Contoh 1
	ahmad := Customer{
		Name:    "Ahmad Maulana Nasution",
		Address: "Bogor, Indonesia",
		Age:     24,
	}
	fmt.Println(ahmad)
	fmt.Println("=============================================")

	// Contoh 2
	hapip := Customer{"Hapip Febriandi", "Bogor, Indonesia", 23} // Harus sesuai dengan posisi dari field/variable yang ada didalam struct
	fmt.Println(hapip)
	fmt.Println("=============================================")
}
