package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (value UserSlice) Len() int {
	return len(value)
}

func (value UserSlice) Less(i, j int) bool {
	return value[i].Age < value[j].Age
}

func (value UserSlice) Swap(i, j int) {
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"Rose", 25},
		{"Jisoo", 27},
		{"Jeni", 26},
		{"Lisa", 25},
	}

	// Sebelum di sorting
	fmt.Println(users)
	sort.Sort(UserSlice(users))
	// Sesudah di sorting
	fmt.Println(users)
}
