package other

import "fmt"

func init() {
	fmt.Println("Init From Other Package")
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}
