package main

import "fmt"

func endApp1() {
	// Recover digunakan untuk menangkap data error dari panic, dan membuat aplikasi bisa berjalan kembali
	// dan di tempatkan pada akhir function, karena kita menggunakan defer yang dimana defer adalah function yang di eksekusi setelah function sebelumnya
	// maka kita tempatkan recover pada function endApp1
	message := recover()
	if message != nil { // Nil merupakan sebuah nilai kosong, kalau di javascript itu kayak NULL
		fmt.Println("Error dengan message", message)
	}
	fmt.Println("End App")
}

func runApp1(error bool) {
	defer endApp1()
	if error {
		panic("Aplikasi Error") // Panic digunakan untuk mengehntikan aplikasi jika terjadi sebuah error, dan defer tetap berjalan walaupun ada error
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp1(true)
}
