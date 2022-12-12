package main

import "fmt"

// Create Type declaration
type Filter func(string) string

// Menjadikan function sebagai parameter
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Aldi", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
