package main

import "fmt"

func sayHelloTo(firstName string, lastName string, age int) {
	fmt.Println("Hello", firstName, lastName, "Age", age)
}

func main() {
	sayHelloTo("Arya", "Rizki Andaru", 22)
}
