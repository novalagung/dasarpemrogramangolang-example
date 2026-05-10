package main

import "fmt"

func main() {
	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "mango", "banana"},
	}

	fmt.Println(data)
}
