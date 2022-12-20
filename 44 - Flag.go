package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Masukkan host")
	var username *string = flag.String("username", "root", "Masukkan user")
	var password *string = flag.String("password", "root", "Masukkan password")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)
}
