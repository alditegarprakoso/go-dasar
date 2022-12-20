package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(1.7)) // dibulatkan ke atas, karena mendekati ke atas
	fmt.Println(math.Round(1.3)) // dibulatkan ke bawah, karena mendekati ke bawah
	fmt.Println(math.Floor(1.7)) // akan dipaksa bulat ke bawah walaupun mendekati ke atas
	fmt.Println(math.Ceil(1.3))  // akan dipaksa bulat ke atas walaupun mendekati ke bawah

	fmt.Println(math.Max(10, 20)) // Mencari nilai terbesar
	fmt.Println(math.Min(10, 20)) // Mencari nilai terkecil
}
