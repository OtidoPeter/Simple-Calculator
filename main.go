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
	outputText("Enter number here: ")
	fmt.Scan(&numOne)

	// fmt.Print("Enter your second number here: ")
	outputText("Enter number here: ")
	fmt.Scan(&numTwo)

	addition, multiplication, division, subtraction := operations(numOne, numTwo)
	// addition := numOne + numTwo

	// multiplication := numOne * numTwo

	// division := numOne / numTwo

	// subtraction := numOne - numTwo

	formattedAdd := fmt.Sprintf("Sum: %.2f\n", addition)
	formattedMult := fmt.Sprintf("Product: %.2f\n", multiplication)
	formattedDiv := fmt.Sprintf("Quotient: %.2f\n", division)
	formattedSub := fmt.Sprintf("Difference: %.2f\n", subtraction)

	/*fmt.Printf(`Sum: %.2f\n

	Product: %.2f\n

	Quotient: %.2f\n

	Difference: %.2f\n`, addition, multiplication, division, subtraction)
	*/
	fmt.Print(formattedAdd, formattedMult, formattedDiv, formattedSub)
}

func outputText(text string) {
	fmt.Print(text)
}

func operations(numOne, numTwo float64) (add float64, mult float64, div float64, subt float64) {
	add = numOne + numTwo
	mult = numOne * numTwo
	div = numOne / numTwo
	subt = numOne - numTwo

	return add, mult, div, subt
}
