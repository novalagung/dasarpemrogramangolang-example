package main

import "fmt"

func main() {
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "wick"
	fmt.Print("nice to meet you ", lastName, "!\n")
	
	const (
		square         = "kotak"
		isToday  bool  = true
		numeric  uint8 = 1
		floatNum       = 2.2
	)

	fmt.Println("===========")
	fmt.Println(square)
	fmt.Println(isToday)
	fmt.Println(floatNum)

	const (
		a = "konstanta"
		b
	)

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("===========")

	const (
		today string = "senin"
		sekarang
		isToday2 = true
	)

	fmt.Println(today)
	fmt.Println(sekarang)
	fmt.Println(isToday)

	fmt.Println("===========")

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	fmt.Println(satu, dua)
	fmt.Println(three, four)
}
