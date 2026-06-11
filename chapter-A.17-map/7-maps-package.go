package main

import (
	"fmt"
	"maps"
)

func main() {
	var chicken = map[string]int{
		"januari":  50,
		"februari": 40,
		"maret":    34,
	}

	// clone map (salinan dangkal)
	var chickenClone = maps.Clone(chicken)
	fmt.Println("clone  :", chickenClone)

	// cek apakah dua map sama persis
	fmt.Println("equal  :", maps.Equal(chicken, chickenClone))

	// hapus semua item yang nilainya di bawah 45
	maps.DeleteFunc(chicken, func(k string, v int) bool {
		return v < 45
	})
	fmt.Println("setelah DeleteFunc:", chicken)

	// salin semua item dari satu map ke map lain
	var extra = map[string]int{"april": 70}
	maps.Copy(chicken, extra)
	fmt.Println("setelah Copy:", chicken)
}
