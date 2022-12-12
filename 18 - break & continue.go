package main

import "fmt"

func main() {
	// Break
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break // break di gunakan untuk menghentikan perulangan, dan tidak akan mengeksekusi sintaks yang berada dibawahnya, maka yang terjadi data yang di print hanya sampai "Perulangan ke 4"
		}
		fmt.Println("Perulangan ke", i)
	}

	fmt.Println("================")

	// Continue
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // continue di gunakan untuk melompati perulangan dengan kondisi yang kita tentukann, dan tidak akan mengeksekusi sintaks yang berada dibawahnya, maka yang terjadi data yang di print hanya sampai bilangan ganjil saja yang tampil
		}
		fmt.Println("Perulangan ke", i)
	}
}
