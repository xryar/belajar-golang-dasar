package main

import "fmt"

func main() {
	name := "Arya"

	if name == "Arya" {
		fmt.Println("Halo Arya!")
	} else if name == "Sugeng" {
		fmt.Println("Halo Sugeng!")
	} else if name == "Joko" {
		fmt.Println("Halo Joko!")
	} else {
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
