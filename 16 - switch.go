package main

import "fmt"

func main() {
	name := "Aldi"

	switch name {
	case "Tegar":
		fmt.Println("Hallo Tegar")
	case "Prakoso":
		fmt.Println("Hallo Prakoso")
	case "Aldi":
		fmt.Println("Hallo Aldi")
	default:
		fmt.Println("Hi, Boleh kenalan ?")
	}

	// Short statename
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama terlalu pendek")
	}

	// Switch Tanpa Expresi
	nilai := 7

	switch {
	case nilai >= 9:
		fmt.Println("Nilai kamu A")
	case nilai <= 8 && nilai >= 6:
		fmt.Println("Nilai kamu B")
	case nilai < 6:
		fmt.Println("Nilai kamu C")
	default:
		fmt.Println("Nilai kamu D")
	}
}
