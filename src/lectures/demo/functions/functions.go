package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println("Hello from Go!")
}

func multiReturn() (int, int, int) {
	return 1, 2, 3
}

func main() {
	greet()

	dozen := double(6)
	fmt.Println("A dozen is", dozen)

	bakersDozen := add(dozen, 1)
	fmt.Println("A baker's dozen is", bakersDozen)

	anotherBakersDozen := add(double(6), 1)
	fmt.Println("Have another", anotherBakersDozen)

	num1, num2, _ := multiReturn()
	fmt.Println(num1, num2)
}
