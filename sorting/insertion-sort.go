package main

import (
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
)

func main() {
	x := utils.GenerateSlice(13)
	fmt.Println(x)

	insertionSort(x)
	fmt.Println(x)
}

func insertionSort(x []int) {
	for i := 1; i < len(x); i++ {
		for j := i - 1; j >= 0; j-- {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
}
