package main

import (
	"fmt"
	"os"
)

func main() {
	// tulis file
	content := []byte("halo\nmari belajar golang\n")
	err := os.WriteFile("/Users/novalagung/Documents/temp/test.txt", content, 0644)
	if err != nil {
		fmt.Println("error tulis:", err)
		return
	}
	fmt.Println("==> file berhasil ditulis")

	// baca file
	data, err := os.ReadFile("/Users/novalagung/Documents/temp/test.txt")
	if err != nil {
		fmt.Println("error baca:", err)
		return
	}
	fmt.Println("==> file berhasil dibaca")
	fmt.Print(string(data))
}
