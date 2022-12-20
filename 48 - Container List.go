package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Aldi")    // Menambahkan data ke paling belakang
	data.PushBack("Tegar")   // Menambahkan data ke paling belakang
	data.PushBack("Prakoso") // Menambahkan data ke paling belakang

	fmt.Println(data.Front().Value)               // Mengambil data dari yang paling depan
	fmt.Println(data.Front().Next().Value)        // Mengambil data setelah data yang paling depan
	fmt.Println(data.Front().Next().Next().Value) // Mengambil data setelah data dari setelah data yang paling depan

	fmt.Println("================================")

	fmt.Println(data.Back().Value)               // Mengambil data dari yang paling belakang
	fmt.Println(data.Back().Prev().Value)        // Mengambil data sebelum data yang paling belakang
	fmt.Println(data.Back().Prev().Prev().Value) // Mengambil data sebelum data yang paling belakang

	fmt.Println("================================")

	data.PushFront("Nia")           // Menambahkan data ke yang paling depan
	fmt.Println(data.Front().Value) // Mengambil data dari yang paling depan

	fmt.Println("================================")

	// Menampilkan Semua Data Dari Depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("================================")

	// Menampilkan Semua Data Dari Depan
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}
}
