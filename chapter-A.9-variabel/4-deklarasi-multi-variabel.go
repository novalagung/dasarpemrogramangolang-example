package main

import "fmt"

func main() {
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	seventh, eighth, ninth := "tujuh", "delapan", "sembilan"

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eighth, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)
}
