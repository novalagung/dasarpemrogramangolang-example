package main

import "fmt"

func main() {
	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "mango"

	fmt.Println(fruits) // [apple mango]
}
