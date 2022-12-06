package main

import "fmt"

func main() {
	var names [3]string // membuat array dengan maksimal 3 data saja bertype data string
	names[0] = "Aldi"
	names[1] = "Tegar"
	names[2] = "Prakoso"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// Sintaks di bawah merupakan array yang diisi secara langsung sekaligus
	var values = [3]int{
		1, 2, 3,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	// len berfungsi untuk menghitung panjang pada array , len ini bukan jumlah data tetapi memang panjang dari si array
	fmt.Println(len(names))
	fmt.Println(len(values))
}
