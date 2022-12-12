package main

import "fmt"

// Membuat function returning multiple values

// Deklarasi function dengan dua return value, yaitu string, dan int
func getBiodata() (string, int) {
	return "Hallo, nama saya Aldi Tegar Prakoso. Saya berumur", 23 // Return value harus sesuai dengan type datanya masing-masing
}

// Deklarasi function dengan 3 return value
func getMyName() (string, string, string) {
	return "Aldi", "Tegar", "Prakoso"
}

func main() {
	fmt.Println(getBiodata())

	fmt.Println("================")

	// Cara dibawah ini adalah cara untuk mengambil return value yang ada pada function, dan harus sesuai dengan jumlah datanya
	// Contoh function getMyName mempunyai 3 return value, dan akan kita ambil dengan cara di bawah ini
	firstName, middleName, lastName := getMyName()
	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)

	fmt.Println("================")

	// Di sintaks di bawah ini, kita akan menghiraukan beberapa return value
	// misalkan kita hanya ingin mengambil return value ke 1 saja dan tidak ingin mengambil return value yang ke 2, dan 3, maka return value ke 2 dan 3 di ganti denga underscore (_)
	myName, _, _ := getMyName()
	fmt.Println(myName)
}
