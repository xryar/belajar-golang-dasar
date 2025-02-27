package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 //copy value dari address1

	// address2.City = "Jakarta"
	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.City = "Jakarta"
	fmt.Println(address1)
	fmt.Println(address2)
}
