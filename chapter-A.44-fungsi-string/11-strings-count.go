package main

import (
	"fmt"
	"strings"
)

func main() {
	var howMany = strings.Count("ethan hunt", "t")
	fmt.Println(howMany) // 2
}
