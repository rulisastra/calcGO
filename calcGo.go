package main

import (
	"fmt"
)

func Calculator() {

	// fmt.Println("running at localhost:2001...")

	// handle panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ran into an error")
			Calculator()
		}
	}()

	functions := map[string]func(int, int) int{
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
		"divide":   divide,
	}

	var currentNumber int
	fmt.Println("Enter first number: ")
	fmt.Scan(&currentNumber)

	// done  := false
	// for !done {
	for true {
		var functionName string
		var number int

		fmt.Println("Choose a function:...\nadd || subtract || multiply || divide ||\ntype \"done\" if you're finished")
		fmt.Scan(&functionName)

		// u should type done when you're done
		if functionName == "done" {
			break
		}

		fmt.Println("Enter the second number:")
		fmt.Scan(&number)

		// current number
		currentNumber = functions[functionName](currentNumber, number)
	}

	fmt.Println("your number is equal to: ")
	fmt.Println(currentNumber)
}

// input output
func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
