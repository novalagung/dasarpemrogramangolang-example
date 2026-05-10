package main

import "fmt"

func main() {
	dst := make([]string, 3)
	src := []string{"watermelon", "pineapple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst) // watermelon pineapple apple
	fmt.Println(src) // watermelon pineapple apple orange
	fmt.Println(n)   // 3

	// ---------------------

	dst = []string{"potato", "potato", "potato"}
	src = []string{"watermelon", "pineapple"}
	n = copy(dst, src)

	fmt.Println(dst) // watermelon pineapple potato
	fmt.Println(src) // watermelon pineapple
	fmt.Println(n)   // 2
}
