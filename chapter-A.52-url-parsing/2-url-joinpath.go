package main

import (
	"fmt"
	"net/url"
)

func main() {
	base := "http://localhost:8080/api"

	joined, err := url.JoinPath(base, "v1", "users", "john wick")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(joined)
	// http://localhost:8080/api/v1/users/john%20wick

	u, _ := url.Parse(base)
	joined2, _ := u.JoinPath("v2", "items")
	fmt.Println(joined2)
	// http://localhost:8080/api/v2/items
}
