package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var customer Customer
	customer.Name = "Arya Rizki Andaru"
	customer.Address = "Jakarta, Indonesia"
	customer.Age = 21

	fmt.Println(customer)
	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)

	// struct literals 1
	joko := Customer{
		Name:    "Joko",
		Address: "Indonesia",
		Age:     30,
	}
	fmt.Println(joko)

	// struct literals 2
	budi := Customer{"Budi", "Indonesia", 20}
	fmt.Println(budi)
}
