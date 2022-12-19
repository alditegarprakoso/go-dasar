package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi Error") // Panic digunakan untuk mengehntikan aplikasi jika terjadi sebuah error, dan defer tetap berjalan walaupun ada error
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
}
