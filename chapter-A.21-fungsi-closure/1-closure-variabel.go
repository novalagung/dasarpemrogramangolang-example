package main

import "fmt"

func main() {
	var getMinMax = func(n []int) (int, int) {
		var minNum, maxNum int
		for i, e := range n {
			switch {
			case i == 0:
				maxNum, minNum = e, e
			case e > maxNum:
				maxNum = e
			case e < minNum:
				minNum = e
			}
		}
		return minNum, maxNum
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}
	var minNum, maxNum = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, minNum, maxNum)
}
