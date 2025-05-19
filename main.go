package main

import "fmt"

func main() {
	var number int
	var otherNumber float64

	fmt.Print("Give the number: ")
	fmt.Scan(&number)

	fmt.Print("Give the other number: ")
	fmt.Scan(&otherNumber)

	add := number + int(otherNumber)
	fmt.Println(add)

	multiply := number * int(otherNumber)
	fmt.Println(multiply)

	division := number / int(otherNumber)
	fmt.Println(division)
}
