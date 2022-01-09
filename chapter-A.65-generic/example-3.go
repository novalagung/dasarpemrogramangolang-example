package main

import (
    "fmt"
)

type UserModel[T int | float64] struct {
	Name string
    Scores []T
}

func (m *UserModel[int|float64]) SetScoresA(scores []int) {
	m.Scores = scores
}

func (m *UserModel[float64]) SetScoresB(scores []float64) {
	m.Scores = scores
}

func main() {
	var m1 UserModel[int]
	m1.Name = "Noval"
	m1.Scores = []int{1, 2, 3}
    fmt.Println("scores:", m1.Scores)

	var m2 UserModel[float64]
	m2.Name = "Noval"
	m2.SetScoresB([]float64{10, 11})
    fmt.Println("scores:", m2.Scores)
}