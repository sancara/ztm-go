//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var color = "Blue"
	var birthYear, age = 1990, 34
	word1, word2, _ := "Hello", "World", "!"
	var (
		firstInitial = "S"
		lastInitial  = "C"
	)
	var ageDays float64
	ageDays = float64(age) * 365.0

	fmt.Println("My favourite color is:", color)
	fmt.Println("My year of birth is", birthYear)
	fmt.Println("I am", age, "years old")
	fmt.Println("compound assignment with ignore _", word1, word2)
	fmt.Println(firstInitial)
	fmt.Println(lastInitial)
	fmt.Println(ageDays)

}
