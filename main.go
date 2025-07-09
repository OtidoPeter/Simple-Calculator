package main

import "fmt"

const appName = " Welcome to the Go - Calculator"
const separator = "-------------------------------"

func main() {
	var numOne float64
	var numTwo float64

	fmt.Println(appName)
	fmt.Println(separator)

	fmt.Println("What operation do you want to carry out?")

	// fmt.Print("Enter your first number here: ")
	numOne = getUserInput("Enter number here: ")
	// fmt.Scan(&numOne)

	// fmt.Print("Enter your second number here: ")
	numTwo = getUserInput("Enter number here: ")
	// fmt.Scan(&numTwo)

	//operations(numOne, numTwo)
	add(numOne, numTwo)
	subtract(numOne, numTwo)
	multiply(numOne, numTwo)
	divide(numOne, numTwo)
	// addition := numOne + numTwo

	// multiplication := numOne * numTwo

	// division := numOne / numTwo

	// subtraction := numOne - numTwo

	/*fmt.Printf(`Sum: %.2f\n

	Product: %.2f\n

	Quotient: %.2f\n

	Difference: %.2f\n`, addition, multiplication, division, subtraction)
	*/
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}

/*
	func operations(numOne, numTwo float64) (float64, float64, float64, float64) {
		addition := numOne + numTwo
		multiplication := numOne * numTwo
		division := numOne / numTwo
		subtraction := numOne - numTwo

		formattedAdd := fmt.Sprintf("Sum: %.2f\n", addition)
		formattedMult := fmt.Sprintf("Product: %.2f\n", multiplication)
		formattedDiv := fmt.Sprintf("Quotient: %.2f\n", division)
		formattedSub := fmt.Sprintf("Difference: %.2f\n", subtraction)

		fmt.Print(formattedAdd, formattedMult, formattedDiv, formattedSub)

		return addition, multiplication, division, subtraction
		// return
	}
*/
func add(numOne, numTwo float64) float64 {
	addition := numOne + numTwo
	reformattedAdd := fmt.Sprintf("Sum: %.1f\n", addition)
	fmt.Print(reformattedAdd)

	return addition
}
func subtract(numOne, numTwo float64) float64 {
	subtraction := numOne - numTwo
	reformattedSubt := fmt.Sprintf("Difference: %.1f\n", subtraction)
	fmt.Print(reformattedSubt)

	return subtraction
}
func multiply(numOne, numTwo float64) float64 {
	multiplication := numOne * numTwo
	reformattedMult := fmt.Sprintf("Product: %.1f\n", multiplication)
	fmt.Print(reformattedMult)

	return multiplication
}
func divide(numOne, numTwo float64) float64 {
	division := numOne / numTwo
	reformattedDiv := fmt.Sprintf("Quotient: %.1f\n", division)
	fmt.Print(reformattedDiv)

	return division
}
