//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import (
	"fmt"
	"math/rand/v2"
)

func greeting(name string) {
	fmt.Println("Hello", name)
}

func displayMessage(msg string) string {
	return msg
}

func add(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func anyNum() int {
	return rand.IntN(100)
}

func any2Num() (int, int) {
	return rand.IntN(100), rand.IntN(100)
}

func main() {
	greeting("Santiago")

	msg := displayMessage("Cosme Fulanito")
	fmt.Println(msg)

	suma := add(2, 3, 4)
	fmt.Println(suma)

	num1 := anyNum()
	fmt.Println(num1)

	num2, num3 := any2Num()
	fmt.Println("num2=", num2, "num3=", num3)

	randomAddition := add(suma, anyNum(), num3)
	fmt.Println("adding=", randomAddition)
}
