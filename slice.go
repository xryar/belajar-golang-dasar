package main

import "fmt"

func main() {
	names := [...]string{
		"Arya",
		"Rizki",
		"Andaru",
		"Joko",
		"Budi",
		"Sugeng",
	}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[0:]
	fmt.Println(slice3)
}