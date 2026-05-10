package main

import "fmt"
import "strings"

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "mango", "banana"}
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")
}
