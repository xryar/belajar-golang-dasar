package main

import "fmt"

func random() any {
	return 100
}

func main() {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Integer", value)
	default:
		fmt.Println("Unknown Type")
	}

}
