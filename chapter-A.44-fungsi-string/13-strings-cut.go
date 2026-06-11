package main

import (
	"fmt"
	"strings"
)

func main() {
	before, after, found := strings.Cut("user:password", ":")
	fmt.Println(before) // user
	fmt.Println(after)  // password
	fmt.Println(found)  // true

	before2, after2, found2 := strings.Cut("nocohere", ":")
	fmt.Println(before2) // nocohere
	fmt.Println(after2)  // (string kosong)
	fmt.Println(found2)  // false
}
