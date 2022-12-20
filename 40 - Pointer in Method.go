package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	aldi := Man{"Aldi"}
	aldi.Married()
	fmt.Println(aldi.Name)
}
