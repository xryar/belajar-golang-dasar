package main

import "fmt"

func main() {
	nilaiAkhir := 90
	nilaiAbsen := 80

	lulusNilaiAkhir := nilaiAkhir > 80
	lulusAbsensi := nilaiAbsen > 80

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}