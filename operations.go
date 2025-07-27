package main

import "fmt"

func operations(numOne, numTwo float64) {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addition := numOne + numTwo
		reformattedAdd := fmt.Sprintf("Sum: %.1f\n", addition)
		fmt.Print(reformattedAdd)
		writeFloatToFile(responseFileName, addition)
	case 2:
		subtraction := numOne - numTwo
		reformattedSubt := fmt.Sprintf("Difference: %.1f\n", subtraction)
		fmt.Print(reformattedSubt)
		writeFloatToFile(responseFileName, subtraction)
	case 3:
		multiplication := numOne * numTwo
		reformattedMult := fmt.Sprintf("Product: %.1f\n", multiplication)
		fmt.Print(reformattedMult)
		writeFloatToFile(responseFileName, multiplication)
	case 4:
		division := numOne / numTwo
		reformattedDiv := fmt.Sprintf("Quotient: %.1f\n", division)
		fmt.Print(reformattedDiv)
		writeFloatToFile(responseFileName, division)
	default:
		fmt.Println("Goodbye! Thank you for choosing our calculator.")
		return
	}
}
