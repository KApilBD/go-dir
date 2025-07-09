package main

import (
	"fmt"
	"strings"
)

func main() {
	var operation string
	var number1, number2 int

	fmt.Println("CALCULATOR GO 0.1")
	fmt.Println("=================")
	fmt.Println("Which operation you want to perform? (add, subtract, multiply, divide)")
	fmt.Scanln(&operation)
	operation = strings.ToLower(operation)

	fmt.Println("Enter first number")
	fmt.Scanln(&number1)

	fmt.Println("Enter second number")
	fmt.Scanln(&number2)

	switch operation {
	case "add":
		fmt.Println("Result:", number1+number2)
	case "subtract":
		fmt.Println("Result:", number1-number2)
	case "multiply":
		fmt.Println("Result:", number1*number2)
	case "divide":
		if number2 == 0 {
			fmt.Println("Error: Cannot divide by zero.")
		} else {
			fmt.Println("Result:", number1/number2)
		}
	default:
		fmt.Println("Invalid operation. Please use add, subtract, multiply, or divide.")
	}
}
