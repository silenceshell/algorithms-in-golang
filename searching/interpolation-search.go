package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	target := x[rand.Intn(len(x))]
	index, err := interpolationSearch(x, target)
	if err == nil {
		fmt.Printf("%d's index is %d", target, index)
	} else {
		fmt.Println(err)
	}
}

func interpolationSearch(x []int, target int) (int, error) {
	left, right := 0, len(x)-1
	if target < x[left] || target > x[right] {
		return -1, fmt.Errorf("not found %d", target)
	}

	min, max := x[left], x[right]
	for left <= right {
		size := right - left
		offset := size * (target - min) / (max - min)
		midd := left + offset
		if x[midd] < target {
			left = midd + 1
			min = x[midd]
		} else if x[midd] > target {
			right = midd - 1
			max = x[midd]
		} else {
			return midd, nil
		}
	}

	return 0, fmt.Errorf("not found %d", target)
}
