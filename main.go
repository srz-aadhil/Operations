package main

import (
	"fmt"
	"operation/addition"
	"operation/division"
	"operation/ispositive"
	"operation/multiplication"
	"operation/subtraction"
)

func main() {
	result_positive := ispositive.Ispositive(-1)
	fmt.Println(result_positive)
	result_add := addition.Add(5, 3)
	fmt.Println(result_add)
	result_subtract := subtraction.Subtract(8, 2)
	fmt.Println(result_subtract)
	result_multiply := multiplication.Multiply(3, 3)
	fmt.Println(result_multiply)
	result_divide, err := division.Divide(5, 0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result_divide)
}
