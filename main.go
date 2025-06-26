package main

import "fmt"

const appName = "Go - Calculator"
const separator = "---------------"

func main() {
	var numOne float64
	var numTwo float64

	fmt.Println(appName)
	fmt.Println(separator)

	// fmt.Print("Enter your first number here: ")
	numOne = getUserInput("Enter number here: ")
	// fmt.Scan(&numOne)

	// fmt.Print("Enter your second number here: ")
	numTwo = getUserInput("Enter number here: ")
	// fmt.Scan(&numTwo)

	operations(numOne, numTwo)
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
