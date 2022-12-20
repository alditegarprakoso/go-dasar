package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Aldi")
	fmt.Println(person)

	var person2 map[string]string = NewMap("")
	if person2 == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person2)
	}
}
