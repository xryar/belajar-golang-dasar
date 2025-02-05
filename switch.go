package main

import "fmt"

func main() {
	name := "Arya"

	switch name {
	case "Arya":
		fmt.Println("Halo Arya")
	case "Budi":
		fmt.Println("Halo Budi")
	case "Joko":
		fmt.Println("Halo Joko")
	case "Sugeng":
		fmt.Println("Halo Sugeng")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// short statement di switch
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
