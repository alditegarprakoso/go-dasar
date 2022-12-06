package main

import "fmt"

func main() {
	type NoKtp string // Menjadikan type alias NoKtp mennjadi type data string
	type Married bool

	var noKtpAldi NoKtp = "01293123129"
	var marriedStatus Married = false
	fmt.Println(noKtpAldi)
	fmt.Println(marriedStatus)
}
