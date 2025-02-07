package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication() {
	// defer adalah function yang dijadwalkan untuk dieksekusi
	// setelah sebuah function selesai dieksekusi
	defer logging()

	fmt.Println("Run Aplication")
}

func main() {
	runApplication()
}
