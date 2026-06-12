package main

import "fmt"

func main() {
	// swap tanpa variabel sementara
	x, y := 10, 20
	x, y = y, x
	fmt.Println(x, y) // 20 10

	// semua RHS dievaluasi sebelum assignment
	i, j := 0, 1
	i, j = i+j, i // i = 0+1 = 1, j = 0 (nilai i lama)
	fmt.Println(i, j) // 1 0
}
