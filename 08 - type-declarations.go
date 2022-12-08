package main

import "fmt"

func main() {
	type NoKtp string // Menjadikan type data alias NoKtp mennjadi type data string
	type Married bool // Menjadikan type data alias Married mennjadi type data boolean (true or false)

	var noKtpAldi NoKtp = "01293123129"
	var marriedStatus Married = false
	fmt.Println(noKtpAldi)
	fmt.Println(marriedStatus)
}
