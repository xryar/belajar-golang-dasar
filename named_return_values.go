package main

import "fmt"

func getCompletedName() (firstName, middleName, lastName string) {
	firstName = "Arya"
	middleName = "Rizki"
	lastName = "Andaru"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getCompletedName()
	fmt.Println("Fullname ", a, b, c)
}
