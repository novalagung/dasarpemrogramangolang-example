package library

import "fmt"

func SayHello(name string) {
	fmt.Println("hello")
	introduce(name)
}

func introduce(name string) {
	fmt.Println("nama saya", name)
}
