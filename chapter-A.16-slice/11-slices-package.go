package main

import (
	"fmt"
	"slices"
)

func main() {
	var fruits = []string{"apple", "grape", "banana", "melon", "grape"}

	// cek apakah elemen ada dalam slice
	fmt.Println(slices.Contains(fruits, "banana")) // true
	fmt.Println(slices.Contains(fruits, "papaya")) // false

	// cari indeks pertama kemunculan elemen
	fmt.Println(slices.Index(fruits, "grape")) // 1

	// urutkan elemen slice (ascending, in-place)
	slices.Sort(fruits)
	fmt.Println("sorted  :", fruits)

	// balik urutan elemen (in-place)
	slices.Reverse(fruits)
	fmt.Println("reversed:", fruits)

	// hapus elemen duplikat yang berurutan
	slices.Sort(fruits)
	fruits = slices.Compact(fruits)
	fmt.Println("compact :", fruits)
}
