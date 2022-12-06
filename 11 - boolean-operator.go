package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	fmt.Println(lulusNilaiAkhir)

	var lulusNilaiAbsen bool = nilaiAbsensi > 80
	fmt.Println(lulusNilaiAbsen)

	var lulus bool = lulusNilaiAkhir && lulusNilaiAbsen
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && nilaiAbsensi >= 80)
}
