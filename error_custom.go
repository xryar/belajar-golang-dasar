package main

import "fmt"

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (nf *notFoundError) Error() string {
	return nf.Message
}

func saveData(id string, data any) error {
	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "arya" {
		return &notFoundError{"data Not Found"}
	}

	return nil
}

func main() {
	err := saveData("arya", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("Validation Error:", validationErr.Error())
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("Not Found Error:", notFoundErr.Error())
		// } else {
		// 	fmt.Println("Unknow Error")
		// }

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("Validation Error:", finalError.Error())
		case *notFoundError:
			fmt.Println("Not Found Error:", finalError.Error())
		default:
			fmt.Println("Unknow Error:", finalError.Error())
		}
	} else {
		fmt.Println("Sukses")
	}
}
