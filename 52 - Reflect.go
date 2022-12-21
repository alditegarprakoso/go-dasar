package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"20"` // Ini adalah struct tag
	Age  int
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Aldi", 23}

	var sampleType reflect.Type = reflect.TypeOf(sample)
	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(1).Name)

	structField := sampleType.Field(0)
	required := structField.Tag.Get("required") // mengambil nilai dari struct tag required yaitu true
	max := structField.Tag.Get("max")           // mengambil nilai dari struct tag max yaitu 20
	fmt.Println(required)
	fmt.Println(max)

	sample.Name = ""
	fmt.Println(IsValid(sample))
}
