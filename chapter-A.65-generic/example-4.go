package main

import "fmt"

type Pair[T any] struct {
	First, Second T
}

type StringPair = Pair[string]

type NumberPair[T int | float64] = Pair[T]

func main() {
	var sp StringPair
	sp.First = "halo"
	sp.Second = "dunia"
	fmt.Println(sp.First, sp.Second)

	var np NumberPair[int]
	np.First = 1
	np.Second = 2
	fmt.Println(np.First, np.Second)

	var fp NumberPair[float64]
	fp.First = 3.14
	fp.Second = 2.71
	fmt.Println(fp.First, fp.Second)
}
