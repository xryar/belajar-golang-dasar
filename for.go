package main

import "fmt"

func main() {

	// for counter := 1; counter <= 10; counter++ {
	// 	fmt.Println("Perulangan ke ", counter)
	// }
	// fmt.Println("Selesai")

	names := []string{"Arya", "Rizki", "Andaru"}
	// manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// otomatis
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
