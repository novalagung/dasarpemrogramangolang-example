package main

import "fmt"

type Person struct {
	name string `tag1`
	age  int    `tag2`
}

type People = Person

func main() {
	var p1 = Person{"wick", 21}
	fmt.Println(p1)

	var p2 = People{"wick", 21}
	fmt.Println(Person(p2))
	fmt.Println(People(p1))
}
