package main

import "fmt"

// skema 1: func(yield func() bool)
// tidak menghasilkan nilai, hanya menandakan satu putaran terjadi
func repeat(n int) func(yield func() bool) {
	return func(yield func() bool) {
		for i := 0; i < n; i++ {
			if !yield() {
				return
			}
		}
	}
}

// skema 2: func(yield func(V) bool)
// menghasilkan satu nilai per iterasi
func counter(n int) func(yield func(int) bool) {
	return func(yield func(int) bool) {
		for i := 1; i <= n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// skema 3: func(yield func(K, V) bool)
// menghasilkan dua nilai per iterasi (key-value)
func enumerate(items []string) func(yield func(int, string) bool) {
	return func(yield func(int, string) bool) {
		for i, v := range items {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	// skema 1: tidak ada nilai, hanya hitung putaran
	count := 0
	for range repeat(5) {
		count++
		fmt.Println("iterasi ke", count)
	}

	fmt.Println("---")

	// skema 2: satu nilai per iterasi
	for v := range counter(5) {
		fmt.Println(v)
	}

	fmt.Println("---")

	// skema 3: dua nilai per iterasi
	fruits := []string{"apple", "mango", "banana"}
	for i, v := range enumerate(fruits) {
		fmt.Println(i, v)
	}

	fmt.Println("---")

	// break menghentikan iterator lebih awal
	for v := range counter(10) {
		if v == 3 {
			break
		}
		fmt.Println(v)
	}
}
