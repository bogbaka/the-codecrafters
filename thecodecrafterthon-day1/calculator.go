package main

import (
	"fmt"
	"strconv"
)

func main() {
	for {

		var qut string
		fmt.Println("Hello lovelies,Welcome to my calculator")
		fmt.Println("Type quit to exits or type continue for the rest of the operations")
		fmt.Scan(&qut)

		if qut == "quit" {
			fmt.Println("Goodbye,my love.")
			break
		} else {
			fmt.Println("continue operations")
			goto start
		}

	start:

		var num1 int
		fmt.Println("enter first number")
		fmt.Scan(&num1)

		var operation string
		fmt.Println("choose an operation: Add, Sub, Mult, Div")
		fmt.Scan(&operation)

		var Index string
		fmt.Println("enter second number")
		fmt.Scan(&Index)
		num2, err := strconv.Atoi(Index)
		if err != nil {
			fmt.Println("only numbers allowed!")
			goto start
		}

		if operation == "+" {
			result := (num1 + num2)
			fmt.Println("The result is: ")
			fmt.Println(result)

		}

		if operation == "-" {
			result := (num1 - num2)
			fmt.Println("The result is: ")
			fmt.Println(result)
		}

		if operation == "*" {
			result := (num1 * num2)
			fmt.Println("The result is: ")
			fmt.Println(result)
		}

		if operation == "/" {
			if num1 == 0 || num2 == 0 {
				fmt.Println("Error: cannot be divided by zero ")
			} else {
				fmt.Println("Result", num1/num2)

			}
		}

	}
}
