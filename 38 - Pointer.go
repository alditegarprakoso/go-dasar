package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// Tidak menggunakan Pointer
	address1 := Address{"Bogor", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Cirebon"

	fmt.Println("Address 1 :", address1) // City dari address1 tidak berubah, karena address2 merupakan duplikat dari address1 yang dimana pas By Value bukan By Reference
	fmt.Println("Address 2 :", address2)

	// Menggunakan Pointer
	address3 := Address{"Bandung", "Jawa Barat", "Indonesia"}
	address4 := &address3 // Ini kita melakukan pointer dengan simbol & untuk mereference data agar terbuhung ke address3

	address4.City = "Jakarta"

	fmt.Println("Address 3 :", address3) // City dari address3 berubah, karena di address4 kita menggunakan pointer yang dimana me-reference ke address3
	fmt.Println("Address 4 :", address4)

	// Menggunakan Pointer dengan penamaan variale yang lengkap
	var address5 Address = Address{"Tangerang", "Jawa Barat", "Indonesia"}
	var address6 *Address = &address5 // Ini kita melakukan pointer dengan simbol & dan * untuk mereference data agar terbuhung ke address5

	address6.City = "Tangerang Selatan"

	fmt.Println("Address 5 :", address5) // City dari address5 berubah, karena di address6 kita menggunakan pointer yang dimana me-reference ke address3
	fmt.Println("Address 6 :", address6)

	// Menggunakan Pointer dengan penamaan variale yang lengkap
	var address7 Address = Address{"Banten", "Jawa Barat", "Indonesia"}
	var address8 *Address = &address7 // Ini kita melakukan pointer dengan simbol & dan * untuk mereference data agar terbuhung ke address5

	address8.City = "Bekasi"

	address8 = &Address{"Sukabumi", "Jawa Barat", "Indonesia"} // Menginisiasi ulang data

	fmt.Println("Address 7 :", address7) // City dari address7 tidak berubah, karena di address8 kita menginisiasi lagi dengan data yang baru
	fmt.Println("Address 8 :", address8)

	var address9 *Address = &address7
	var address10 *Address = &address7

	*address10 = Address{"Malang", "Jawa Timur", "Indonesia"} // Mereferensi semua data yang menggunakan Struct Address
	fmt.Println("Address 7 :", address7)
	fmt.Println("Address 9 :", address9)
	fmt.Println("Address 10 :", address10)

	var alamat1 *Address = new(Address) // Ini adalah new Function, yang berfungisi untuk membuat sebuah reference baru dengan nilai yang kosong
	alamat1.City = "Kota Bogor"
	fmt.Println("Alamat 1:", alamat1)
}
