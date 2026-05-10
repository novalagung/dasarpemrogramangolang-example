package main

import "fmt"

func main() {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "mango", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
}
