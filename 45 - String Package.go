package main

import (
	"fmt"
	"strings"
)

func main() {
	// Contains
	fmt.Println(strings.Contains("Aldi Tegar Prakoso", "Aldi")) // Untuk mengecek apakah ada kata Aldi di dalam string Aldi Tegar Prakoso
	fmt.Println(strings.Contains("Nia Sarniati", "Aldi"))       // Untuk mengecek apakah ada kata Aldi di dalam string Nia Sarniati

	// Split
	fmt.Println(strings.Split("Aldi Tegar Prakoso", " "))

	// To Lower, To Upper, To Title
	fmt.Println(strings.ToLower("Aldi Tegar Prakoso"))
	fmt.Println(strings.ToUpper("Aldi Tegar Prakoso"))
	fmt.Println(strings.ToTitle("aldi tegar prakoso"))

	// Trim
	fmt.Println(strings.Trim("             Aldi Tegar Prakoso         ", " ")) // Untuk menghapus spasi diposisi kiri dan kanan dalam string Aldi Tegar Prakoso

	// Replace All
	fmt.Println(strings.ReplaceAll("Aldi Aldi Aldi", "Aldi", "Nia"))
}
