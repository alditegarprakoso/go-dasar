package main

import (
	"fmt"
	"os"
)

func main() {
	// Kalau mau run kaya gini "go run nama_file.go Aldi Tegar Prakoso", karena ini mau ambil argument dari package OS
	args := os.Args // Args digunakan untuk mengambil argument
	fmt.Println(args)

	// fmt.Println(args[1])
	// fmt.Println(args[2])
	// fmt.Println(args[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname:", name)
	} else {
		fmt.Println("Error:", err)
	}

	// Untuk mengambil environment
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

}
