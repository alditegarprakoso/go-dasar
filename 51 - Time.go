package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()             // Inisialisasi Time
	fmt.Println(now)              // Waktu saat ini
	fmt.Println(now.Year())       // Ambil Tahun
	fmt.Println(now.Month())      // Ambil Bulan
	fmt.Println(now.Day())        // Ambil Tanggal
	fmt.Println(now.Hour())       // Ambil jam
	fmt.Println(now.Minute())     // Ambil menit
	fmt.Println(now.Second())     // Ambil detik
	fmt.Println(now.Nanosecond()) // Ambil nanosecond

	// Membuat waktu
	utc := time.Date(2022, 9, 27, 23, 59, 59, 59, time.UTC) // Urutan yang ada di parameter => tahun, bulan, tanggal, jam, menit, detik, nanosecond, dan lokasi
	fmt.Println(utc)

	layout := "2006-01-02" // Jika ingin membuat sebuat tanggal, maka layout harus seperti ini
	parse, err := time.Parse(layout, "1999-09-27")
	if err == nil {
		fmt.Println(parse)
	} else {
		fmt.Println(err.Error())
	}

	// Atau bisa seperti ini
	parse2, err2 := time.Parse(time.RFC3339, "1999-09-27T23:59:59Z")
	if err2 == nil {
		fmt.Println(parse2)
	} else {
		fmt.Println(err2.Error())
	}
}
