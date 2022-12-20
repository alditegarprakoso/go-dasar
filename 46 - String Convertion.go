package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true") // Mengkonversi dari string ke type data boolean, dan value harus sesuai dengan type data yang di konversi
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64) // Mengkonversi dari string ke type data int, dan value harus sesuai dengan type data yang di konversi
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 10) // Konversi dari int ke string
	fmt.Println(value)

	valueString, err := strconv.Atoi("2000000") // Konversi dari string to int menggunakan Atoi
	if err == nil {
		fmt.Println(valueString)
	} else {
		fmt.Println(err.Error())
	}

	valueInt := strconv.Itoa(2000000) // Konversi dari int to string menggunakan Itoa
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err.Error())
	}
}
