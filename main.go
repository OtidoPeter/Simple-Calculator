package main

import "fmt"

func main() {
	const APP_NAME = "GoCalc 1.0"
	const SEPARATOR = "------------------"
	var num1 float64
	var num2 float64

	fmt.Print("Give the number: ")
	fmt.Scan(&num1)

	fmt.Print("Give the other number: ")
	fmt.Scan(&num2)

	fmt.Println(APP_NAME)

	add := num1 + num2
	fmt.Printf("add: %.2f\n", add)
	fmt.Println(SEPARATOR)

	multiply := num1 * num2
	fmt.Printf("multiply: %.2f\n", multiply)
	fmt.Println(SEPARATOR)

	division := num1 / num2
	fmt.Printf("division: %.2f\n", division)
	fmt.Println(SEPARATOR)
}
