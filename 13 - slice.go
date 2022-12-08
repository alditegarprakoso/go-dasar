package main

import "fmt"

func main() {
	var months = [...]string{ // [...] ini berfungsi untuk memberikan panjang array secara otomatis
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) // Untuk mengetahui panjang dari slice
	fmt.Println(cap(slice1)) // Untuk mengetahui kapasitas dari slice

	months[5] = "juniDiUbah" // Jika kita ubah sebuah data pada arraynya maka slicenya pun akan berubah karena slice terkoneksi dengan array secara langsung
	fmt.Println(slice1)

	slice1[0] = "meiDiUbah" // Jika kita juga mengubah data yang ada pada slice maka data pada array pun akan ikut berubah karena memang array dan slice itu saling terhubung
	fmt.Println(months)

	var days = [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	var slice2 = days[5:] // Kita slice mulai dari posisi 6 sampe akhir
	fmt.Println(slice2)

	var slice3 = append(slice2, "nambahHari") // Append digunakan untuk menambah data baru	ke array di paling akhir, dan jika lengthnya tidak cukup, maka dia akan membuat array yang baru yang tidak terkoneksi dengan array sebelumnya
	fmt.Println(slice3)
	slice3[1] = "mingguDiEdit"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 3) // fungsi make ini berfungsi untuk membuat slice baru

	newSlice[0] = "Aldi"
	newSlice[1] = "Tegar"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice) // copy ini berfungsi untuk mengcopy slice yang sudah ada, dan pastikan ya lenght/len, kapasitas nya harus sama, kalau tidak nanti bakal kepotong
	fmt.Println(copySlice)
}
