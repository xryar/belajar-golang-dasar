package main

import "fmt"

func main() {
	// var name string
	// bisa pake ini :=
	name := "Arya Rizki"
	fmt.Println(name)

	name = "Arya Andaru"
	fmt.Println(name)

	// multiple variable
	var(
		firstName = "Arya"
		middleName = "Rizki"
		lastName = "Andaru"
	)

	fmt.Println(firstName)
	fmt.Println(middleName)
	fmt.Println(lastName)
}