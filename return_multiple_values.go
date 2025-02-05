package main

import "fmt"

func getFullname() (string, string) {
	return "Arya", "Rizki"
}

func main() {
	// firstName, lastName := getFullname()
	firstName, _ := getFullname()
	fmt.Println(firstName)
}
