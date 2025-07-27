package main

import (
	"fmt"

	"example.com/Simple-Calculator/fileops"
)

func operations(numOne, numTwo float64) {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addition := numOne + numTwo
		reformattedAdd := fmt.Sprintf("Sum: %.1f\n", addition)
		fmt.Print(reformattedAdd)
		fileops.WriteFloatToFile(responseFileName, addition)
	case 2:
		subtraction := numOne - numTwo
		reformattedSubt := fmt.Sprintf("Difference: %.1f\n", subtraction)
		fmt.Print(reformattedSubt)
		fileops.WriteFloatToFile(responseFileName, subtraction)
	case 3:
		multiplication := numOne * numTwo
		reformattedMult := fmt.Sprintf("Product: %.1f\n", multiplication)
		fmt.Print(reformattedMult)
		fileops.WriteFloatToFile(responseFileName, multiplication)
	case 4:
		division := numOne / numTwo
		reformattedDiv := fmt.Sprintf("Quotient: %.1f\n", division)
		fmt.Print(reformattedDiv)
		fileops.WriteFloatToFile(responseFileName, division)
	default:
		fmt.Println("Goodbye! Thank you for choosing our calculator.")
		return
	}
}
