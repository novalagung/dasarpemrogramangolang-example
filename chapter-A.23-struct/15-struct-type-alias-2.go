package main

import "fmt"

type People1 struct {
	name string
	age int
}

type People2 = struct {
	name string
	age int
}

func main() {
	people := People1{"wick", 21}
	fmt.Println(People1(people))

	person := People2{"wick", 21}
	fmt.Println(People2(person))
}
