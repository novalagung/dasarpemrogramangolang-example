package main

import "fmt"
import "regexp"

func main() {
	var text = "banana burger soup"
	var regex, _ = regexp.Compile(`[a-b]+`) // split dengan separator adalah karakter "a" dan/atau "b"

	var str = regex.Split(text, -1)
	fmt.Printf("%#v \n", str)
	// []string{"", "n", "n", " ", "urger soup"}
}
