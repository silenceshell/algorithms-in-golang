package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var items []int = []int{1, 3, 22, 32, 40, 123, 224, 344, 555}
	var targets []int = make([]int, 0, 5)

	for _, v := range items {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	targets = append(targets, items[rand.Intn(len(items))], -1, 4, 33, 40, 125, 500, 600)
	for _, target := range targets {
		index, err := binarySearch(items, target)
		if err == nil {
			fmt.Printf("%d's index is %d\r\n", target, index)
		} else {
			fmt.Printf("%d not found\r\n", target)
		}
	}
}

func binarySearch(items []int, target int) (int, error) {
	left, right := 0, len(items)-1

	for left <= right {
		var mid = left + (right-left)/2
		if items[mid] < target {
			left = mid + 1
		} else if items[mid] > target {
			right = mid - 1
		} else {
			return mid, nil
		}
	}

	return 0, fmt.Errorf("%d not found", target)
}
