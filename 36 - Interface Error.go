package main

import (
	"errors"
	"fmt"
)

func Pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil // karena tidak ada error, maka bisa return yang ke 2 bisa menggunakan nil
	}
}

func main() {
	// Membuat Error
	// errors itu sudah ada bawaan dari golang, yang di mana dia adalah sebuah interface yang didalamnya terdapar function Error()
	var contohError error = errors.New("Ups Error")
	fmt.Println(contohError.Error())

	hasil, err := Pembagi(100, 10) // variable hasil dan err adalah merupakan return dari si function Pembagi()
	if err == nil {
		fmt.Println("Hasil:", hasil)
	} else {
		fmt.Println("Error:", err.Error())
	}
}
