package main

import (
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
)

func main() {
	x := utils.GenerateSlice(13)
	fmt.Println(x)

	selectionSort(x)
	fmt.Println(x)
}

func selectionSort(x []int) {
	for i := 0; i < len(x); i++ {
		var index int = i
		for j := i; j < len(x); j++ {
			if x[j] < x[index] {
				index = j
			}
		}
		if x[index] < x[i] {
			x[i], x[index] = x[index], x[i]
		}
	}
}