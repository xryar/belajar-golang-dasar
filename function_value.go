package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodBye := getGoodBye
	contoh := getGoodBye

	fmt.Println(goodBye("Arya"))
	fmt.Println(contoh("Arya"))
}
