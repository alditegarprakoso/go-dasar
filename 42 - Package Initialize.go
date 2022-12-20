package main

import (
	"fmt"
	"go-dasar/database"

	// underscore yang ada di depan package merupakan blank inditifier
	// yang di mana digunakan jika ingin import package tapi tidak digunakan sama sekali dan hanya untuk menjalankan function ini yang ada di dalamnya
	_ "go-dasar/other"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)

}
