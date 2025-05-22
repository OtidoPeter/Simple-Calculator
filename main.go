package main

import "fmt"

func main() {
	const APP_NAME = "GoCalc 1.0"
	const SEPARATOR = "------------------"
	var num1 float64
	var num2 float64
	num1 = getUserInput("Give the number: ")
	//fmt.Print("Give the number: ")
	// fmt.Scan(&num1)

	num2 = getUserInput("Give the other number: ")
	// fmt.Print("Give the other number: ")
	// fmt.Scan(&num2)

	fmt.Println(SEPARATOR)
	fmt.Println(APP_NAME)
	fmt.Println(SEPARATOR)

	add, multiply, division, subtraction := operations(num1, num2)
	fmt.Printf("Add: %.2f\n", add)
	fmt.Println(SEPARATOR)
	fmt.Printf("Multiply: %.2f\n", multiply)
	fmt.Println(SEPARATOR)
	fmt.Printf("Division: %.2f\n", division)
	fmt.Println(SEPARATOR)
	fmt.Printf("Subtraction: %.2f\n", subtraction)
	fmt.Println(SEPARATOR)

	/*add := num1 + num2
	fmt.Printf("Add: %.2f\n", add)
	fmt.Println(SEPARATOR)

	multiply := num1 * num2
	fmt.Printf("Multiply: %.2f\n", multiply)
	fmt.Println(SEPARATOR)

	division := num1 / num2
	fmt.Printf("Division: %.2f\n", division)
	fmt.Println(SEPARATOR)

	subtraction := num1 - num2
	fmt.Printf("Subtraction: %.2f\n", subtraction)
	fmt.Println(SEPARATOR)*/
}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func operations(num1, num2 float64) (float64, float64, float64, float64) {

	operationsMultiply := num1 * num2
	operationsAdd := num1 + num2
	operationsDiv := num1 / num2
	operationsSubt := num1 - num2
	return operationsMultiply, operationsAdd, operationsDiv, operationsSubt

}
