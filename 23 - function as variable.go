package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	// Disini kita akan menempatkan variable sayGoodBye dengan isian dari function getGoodBye
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("Aldi"))
}
