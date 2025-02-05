package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Arya"
	names[1] = "Rizki"
	names[2] = "Andaru"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	values :=  [...]int{
		90,
		80,
		70,
		100,
		110,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}