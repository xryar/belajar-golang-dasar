package main

import "fmt"

func main() {
	// closure adalah kemampuan sebuah function yang dapat berinteraksi
	// dengan data-data yang ada di sekitarnya dalam scope yang sama
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println(counter)
}
