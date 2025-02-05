package main

import "fmt"

func main() {
	a := 10
	b := 10
	a1 := 8
	a2 := 5

	c := a + b
	fmt.Println(c)

	c += 10
	fmt.Println(c)

	d := a1 % a2
	fmt.Println(d)

	d += 7
	fmt.Println(d)

	j := 1
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)

}