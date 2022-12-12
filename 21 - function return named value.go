package main

import "fmt"

// Disini kita bisa menambahkan nama/variable untuk return value pada function
// Contohnya sebagai di bawah ini

func getFullName() (firstName string, middleName string, lastName string) { // Inisiasi nama/variable return value
	firstName = "Aldi"
	middleName = "Tegar"
	lastName = "Prakoso"

	return // jika sudah menggunakan nama/variable return value, maka tidak usah dipanggil kembali nama/variablenya, cukup dengen "return" saja
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
