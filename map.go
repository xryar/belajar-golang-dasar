package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Arya",
		"address": "Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := map[string]string{
		"title":  "Buku Go-Lang",
		"author": "Eko Kurniawan",
		"wrong":  "Ups",
	}
	fmt.Println(book)
	delete(book, "wrong")
	fmt.Println(book)
}
