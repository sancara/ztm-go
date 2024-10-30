package main

import "fmt"

func main() {
	// compound assignment:
	word1, word2, _ := "Hello", "World", "!"
	// by using the underscore the exclamation mark is ignored
	// fmt.Println() -> Hello World

	fmt.Println(word1, word2)
}
