package main

import (
	"fmt"
	"strings"
)

func main() {
	after, found := strings.CutPrefix("foobar", "foo")
	fmt.Println(after) // bar
	fmt.Println(found) // true

	before, found2 := strings.CutSuffix("foobar", "bar")
	fmt.Println(before) // foo
	fmt.Println(found2) // true
}
