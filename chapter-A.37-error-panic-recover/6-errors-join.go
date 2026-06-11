package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("database connection failed")
	err2 := errors.New("cache unavailable")

	combined := errors.Join(err1, err2)
	fmt.Println(combined)
}
