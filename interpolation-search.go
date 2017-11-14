package main

import "fmt"

func main() {
	x := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	target := 63
	index, err := interpolationSearch(x, target)
	if err == nil {
		fmt.Printf("%d's index is %d", target, index)
	}
}

func interpolationSearch(x []int, target int) (int, error) {
	left, right := 0, len(x)-1
	min, max := x[left], x[right]

	for left <= right {
		size := right - left
		inter := (size -1) * (target - min) / (max - min)

	}
}