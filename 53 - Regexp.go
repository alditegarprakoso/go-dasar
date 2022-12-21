package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("a([a-z])([a-z])i") // Regex untuk huruf yang di awali oleh huruf "a" dan diakhiri dengan huruf "i"

	fmt.Println(regex.MatchString("aldi"))
	fmt.Println(regex.MatchString("aldo"))
	fmt.Println(regex.MatchString("aLDi"))

	// ini kaya ngefilter gitu pokoknya, sesuai dengan regex yang sudah kita buat sebelumnya
	// angka 3 itu adalah jumlah data yang ingin diambil dari hasil pencariannya
	// jika ingin mencari semua datanya, angka 3 bisa diubah menjadi -1
	search := regex.FindAllString("aldi aldo alfi ali aliando", 3)
	fmt.Println(search)
}
