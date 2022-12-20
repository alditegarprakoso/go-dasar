package main

import "fmt"

type Alamat struct {
	City, Province, Country string
}

func ChangeCountryToIndonesiaWithoutPointer(alamat Alamat) {
	alamat.Country = "Indonesia"
}

func ChangeCountryToIndonesiaWithPointer(alamat *Alamat) {
	alamat.Country = "Indonesia"
}

func main() {
	var alamat = Alamat{
		City:     "Bogor",
		Province: "Jawa Barat",
		Country:  "",
	}

	ChangeCountryToIndonesiaWithoutPointer(alamat) // Melakukan perubahan country dengan function yang tidak menggunakan pointer
	fmt.Println(alamat)                            // Tidak akan berubah, karena data yang dikirim kepada function akan dikirim menjadi duplikat, sehingga tidak mempengaruhi data master

	ChangeCountryToIndonesiaWithPointer(&alamat) // Melakukan perubahan country dengan function yang menggunakan pointer
	fmt.Println(alamat)                          // Akan berubah, karena data yang dikirim kepada function akan dikirim menjadi reference yang merujuk pada data master
}
