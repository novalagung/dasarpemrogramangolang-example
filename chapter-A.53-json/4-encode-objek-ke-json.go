package main

import "encoding/json"
import "fmt"

type User struct {
	FullName string `json:"Name"`
	Age      int
}

func main() {
	var object = []User{{FullName: "john wick", Age: 27}, {FullName: "ethan hunt", Age: 32}}
	var jsonData, err = json.Marshal(object)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)
	fmt.Println(jsonString)
}
