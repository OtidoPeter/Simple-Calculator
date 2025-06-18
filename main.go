package main

import "fmt"

func main() {
	const appName = "Go - Calculator"
	const separator = "---------------"
	var numOne float64
	var numTwo float64

	fmt.Println(appName)
	fmt.Println(separator)

	fmt.Print("Enter your first number here: ")
	fmt.Scan(&numOne)

	fmt.Print("Enter your second number here: ")
	fmt.Scan(&numTwo)

	addition := numOne + numTwo
	fmt.Println(addition)

	multiplication := numOne * numTwo
	fmt.Println(multiplication)

	division := numOne / numTwo
	fmt.Println(division)

	subtraction := numOne - numTwo
	fmt.Println(subtraction)
}
