package main

import (
	"belajar-golang-level-akses-struct/library"
	"fmt"
)

func main() {
	var s1 = library.Student{"ethan", 21}
	fmt.Println("name ", s1.Name)
	fmt.Println("grade", s1.Grade)
}
