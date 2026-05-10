package main

import "encoding/base64"
import "fmt"

func main() {
	var data = "http://developer.com/"

	var encodedString = base64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(encodedString)

	var decodedByte, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println(decodedString)
}
