package main

import "fmt"

func main() {
	data := []string{"superman", "aquaman", "wonder woman"}

	for _, each := range data {

		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic occurred on looping", each, "| message:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("some error happened")
		}()

	}
}
