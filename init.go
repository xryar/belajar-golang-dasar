package main

import (
	"fmt"
	"hello-world/database"
	_ "hello-world/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
