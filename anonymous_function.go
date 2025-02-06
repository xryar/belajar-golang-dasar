package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You have Blocked", name)
		return
	} else {
		fmt.Println("Welcome", name)
		return
	}
}

func main() {
	// bisa seperti ini
	blacklist := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Arya", blacklist)

	// atau seperti ini
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}
