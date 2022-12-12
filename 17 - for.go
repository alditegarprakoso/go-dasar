package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan counter ke : ", counter)
		counter++
	}

	fmt.Println("=============================")

	// For dengan statement
	for count := 1; count <= 10; count++ {
		fmt.Println("Perulangan count ke : ", count)
	}

	fmt.Println("=============================")

	// For loop data Slice
	slice := []string{"Aldi", "Tegar", "Prakoso"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	fmt.Println("=============================")

	// For with range
	for i, value := range slice {
		fmt.Println("Index ke", i, "adalah", value)
	}

	fmt.Println("=============================")

	// Jika hanya ingin valuenya saja, tidak butuh indexnya, maka index atau i diganti menjadi underscore (_)
	for _, value := range slice {
		fmt.Println(value)
	}

	fmt.Println("=============================")

	// For loop type data map
	person := make(map[string]string)
	person["name"] = "Aldi Tegar Prakoso"
	person["title"] = "Fullstack Developer"

	for key, value := range person { // karena ini adalah map, maka gunakan key
		fmt.Println(key, "=", value)
	}
}
