package main

import "fmt"

func main() {
	fmt.Println(min(3, 7, 1, 5)) // 1
	fmt.Println(max(3, 7, 1, 5)) // 7

	var a, b = 10, 20
	fmt.Println(min(a, b)) // 10
	fmt.Println(max(a, b)) // 20

	fmt.Println(min("banana", "apple", "cherry")) // apple
	fmt.Println(max("banana", "apple", "cherry")) // cherry
}
