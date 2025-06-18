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

	multiplication := numOne * numTwo

	division := numOne / numTwo

	subtraction := numOne - numTwo

	fmt.Printf("Sum: %.2f\nProduct: %.2f\nQuotient: %.2f\nDifference: %.2f\n", addition, multiplication, division, subtraction)
}
