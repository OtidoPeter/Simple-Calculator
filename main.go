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

	fmt.Println(SEPARATOR)
	fmt.Println(APP_NAME)
	fmt.Println(SEPARATOR)

	add := num1 + num2
	fmt.Printf("Add: %.2f\n", add)
	fmt.Println(SEPARATOR)

	multiply := num1 * num2
	fmt.Printf("Multiply: %.2f\n", multiply)
	fmt.Println(SEPARATOR)

	division := num1 / num2
	fmt.Printf("Division: %.2f\n", division)
	fmt.Println(SEPARATOR)

	subtraction := num1 - num2
	fmt.Printf("Subtraction: %.2f\n", subtraction)
	fmt.Println(SEPARATOR)
}
