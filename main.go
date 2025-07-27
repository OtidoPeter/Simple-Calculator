package main

import (
	"errors"
	"fmt"
)

const responseFileName = "response.txt"

func main() {
	var numOne float64
	var numTwo float64

	const appName = "Go Calculator"
	const separator = "-----------------------"

	fmt.Println(appName)
	fmt.Println(separator)

	numOne, err := getUserInput("Enter number here: ")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Scan(&numOne)

	numTwo, err = getUserInput("Enter number here: ")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Scan(&numTwo)

	//fmt.Printf("Sum: %.1f\nDifference: %.1f\nProduct: %.1f\nQuotient: %.1f\n", addition, subtraction, multiplication, division)
	operations(numOne, numTwo)

}
func getUserInput(text string) (float64, error) {
	var userInput float64
	fmt.Print(text)
	_, err := fmt.Scan(&userInput)

	if err != nil {
		return 0, errors.New("Invalid input! Please enter a float.")
	}

	return userInput, nil
}
