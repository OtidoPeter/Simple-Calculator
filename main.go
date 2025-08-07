package main

import (
	"errors"
	"fmt"

	"example.com/Simple-Calculator/fileops"
)

const resultsFileName = "results.txt"

func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		return 0, errors.New("Only accepts numbers")
	}

	return userInput, nil
}

func main() {
	const appName = "Go Calculator V-12"
	const separator = "-------------------"

	fmt.Println(appName)
	fmt.Println(separator)

	operations()

	var choice int
	fmt.Print("Your choice: ")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		numOne, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numOne)
		if err != nil {
			fmt.Println(err)
		}

		numTwo, err := getUserInput("Enter number here: ")
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Scan(&numTwo)

		addition := numOne + numTwo
		//fmt.Printf("Sum: %.1f\n", addition)
		reformattedAdd := fmt.Sprintf("SUM: %.1f\n", addition)
		fmt.Print(reformattedAdd)
		fileops.WriteFloatToFile(resultsFileName, addition)
		return
	case 2:
		numOne, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numOne)
		if err != nil {
			fmt.Println(err)
		}

		numTwo, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numTwo)
		if err != nil {
			fmt.Println(err)
		}

		subtraction := numOne - numTwo
		//fmt.Printf("DIFFERENCE: %.1f\n", subtraction)
		reformattedSubt := fmt.Sprintf("DIFFERENCE: %.1f\n", subtraction)
		fmt.Print(reformattedSubt)
		fileops.WriteFloatToFile(resultsFileName, subtraction)
		return
	case 3:
		numOne, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numOne)
		if err != nil {
			fmt.Println(err)
		}

		numTwo, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numTwo)
		if err != nil {
			fmt.Println(err)
		}

		multiplication := numOne * numTwo
		//fmt.Printf("PRODUCT: %.1f\n", multiplication)
		reformattedMult := fmt.Sprintf("PRODUCT: %.1f\n", multiplication)
		fmt.Print(reformattedMult)
		fileops.WriteFloatToFile(resultsFileName, multiplication)
		return
	case 4:
		numOne, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numOne)
		if err != nil {
			fmt.Println(err)
		}

		numTwo, err := getUserInput("Enter number here: ")
		//fmt.Scan(&numTwo)
		if err != nil {
			fmt.Println(err)
		}

		division := numOne / numTwo
		//fmt.Printf("QUOTIENT: %.1f\n", division)
		reformattedDiv := fmt.Sprintf("QUOTIENT: %.1f\n", division)
		fmt.Print(reformattedDiv)
		fileops.WriteFloatToFile(resultsFileName, division)
		return
	case 5:
		lastResult, err := fileops.GetFloatFromFile(resultsFileName)
		if err != nil {
			fmt.Println("ERROR")
			fmt.Println(err)
			fmt.Println("---------------")
		}
		fmt.Println("Last saved result:", lastResult)
		return
	default:
		fmt.Println("Thank you!")
		return
	}
}
