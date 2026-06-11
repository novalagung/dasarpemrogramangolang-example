package main

import (
	"fmt"
	"strings"
)

func main() {
	var text = "banana"
	var find = "a"
	var replaceWith = "o"

	var result = strings.ReplaceAll(text, find, replaceWith)
	fmt.Println(result) // bonono
}
