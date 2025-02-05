package main

import "fmt"

func main() {
	type NoKTP string

	var ktpArya NoKTP = "11111111"

	var contoh string = "2222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpArya)
	fmt.Println(contohKtp)
}