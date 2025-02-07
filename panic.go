package main

import "fmt"

func endApp() {
	fmt.Println("End Program")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

// panic adalah fitur yang menghentikan
// semua operasi yang ada di golang saat itu juga
func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}

	// implemetasi recover yang salah
	// message := recover()
	// fmt.Println("Terjadi error", message)
	// cara yang benar ialah mengimplementasikannya di function yang di defer
}

func main() {
	runApp(true)
}
