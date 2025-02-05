package main

import "fmt"

func main() {
	genap := []int{}
	ganjil := []int{}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			genap = append(genap, i)
			continue
		} else {
			ganjil = append(ganjil, i)
			fmt.Println("Perulangan ke", i)
			continue
		}
	}

	fmt.Println("Bilangan genap", genap)
	fmt.Println("Bilangan ganjil", ganjil)
}
