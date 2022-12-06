package main

import "fmt"

func main() {
	var name1 = "Aldi"
	var name2 = "Aldi"
	var name3 = "Tegar"

	var result bool = name1 == name2
	fmt.Println(result)

	result = name2 == name3
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
