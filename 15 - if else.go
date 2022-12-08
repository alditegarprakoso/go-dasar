package main

import "fmt"

func main() {
	name := "Tegar"
	secondName := "Prakoso"

	if name == "Aldi" {
		fmt.Println("Hello Aldi")
	} else {
		fmt.Println("Hi, Boleh kenalan ?")
	}

	if name == "Tegar" {
		fmt.Println("Anda Bukan Aldi")
	}

	if secondName == "Tegar" {
		fmt.Println("Hello Prakoso")
	} else {
		fmt.Println("Hi, Boleh kenalan ?")
	}

	if name == "Aldi" {
		fmt.Println("Hello Aldi")
	} else if name == "Prakoso" {
		fmt.Println("Hello Prakoso")
	} else {
		fmt.Println("Hello Tegar, boleh kenalan ?")
	}

	// Short statement
	if length := len(secondName); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
