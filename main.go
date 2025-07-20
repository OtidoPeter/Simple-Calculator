package main

import (
	"fmt"
	"os"
)

const responseFileName = "response.txt"

func writeResponseToFile(response float64) {
	responseText := fmt.Sprint(response)
	os.WriteFile(responseFileName, []byte(responseText), 0644)
}

func main() {
	var numOne float64
	var numTwo float64

	const appName = "Go Calculator"
	const separator = "-----------------------"

	fmt.Println(appName)
	fmt.Println(separator)

	numOne = getUserInput("Enter number here: ")
	//fmt.Scan(&numOne)

	numTwo = getUserInput("Enter number here: ")
	//fmt.Scan(&numTwo)

	//fmt.Printf("Sum: %.1f\nDifference: %.1f\nProduct: %.1f\nQuotient: %.1f\n", addition, subtraction, multiplication, division)
	operations(numOne, numTwo)

}
func getUserInput(text string) float64 {
	var userInput float64
	fmt.Print(text)
	fmt.Scan(&userInput)

	return userInput
}

func operations(numOne, numTwo float64) {
	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		addition := numOne + numTwo
		reformattedAdd := fmt.Sprintf("Sum: %.1f\n", addition)
		fmt.Print(reformattedAdd)
		writeResponseToFile(addition)
	case 2:
		subtraction := numOne - numTwo
		reformattedSubt := fmt.Sprintf("Difference: %.1f\n", subtraction)
		fmt.Print(reformattedSubt)
		writeResponseToFile(subtraction)
	case 3:
		multiplication := numOne * numTwo
		reformattedMult := fmt.Sprintf("Product: %.1f\n", multiplication)
		fmt.Print(reformattedMult)
		writeResponseToFile(multiplication)
	case 4:
		division := numOne / numTwo
		reformattedDiv := fmt.Sprintf("Quotient: %.1f\n", division)
		fmt.Print(reformattedDiv)
		writeResponseToFile(division)
	default:
		fmt.Println("Goodbye! Thank you for choosing our calculator.")
		return
	}

}
