package main

import "fmt"

func main() {

	// map[keyTypeData]valueTypeData -> ini adalah struktur dari map, jadi di map ada yang namanya Key dan Value, mereka mempunyai type data masing-masing
	person := map[string]string{
		"name":    "Aldi Tegar Prakoso",
		"address": "Bogor",
	}

	person["title"] = "Programmer"   // Sintaks ini berfungsi untuk menambahkan key dan value baru pada person
	person["phone"] = "083454356543" // Sintaks ini berfungsi untuk menambahkan key dan value baru pada person

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	fmt.Println(person["phone"])
	fmt.Println(len(person))

	delete(person, "phone") // Fungsi untuk menghapus data phone pada person

	fmt.Println("\n") // New Line
	fmt.Println("================================")
	fmt.Println("\n")

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])
	fmt.Println(person["phone"]) // Tidak akan muncul karena sudah di hapus
	fmt.Println(len(person))
}
