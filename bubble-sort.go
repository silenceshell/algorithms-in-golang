package main

import (
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
)

func main() {
	x := utils.GenerateSlice(13)
	fmt.Println(x)

	bubbleSort(x)
	fmt.Println(x)
}

func bubbleSort(x []int) {
	for i:=0; i < len(x); i++ {
		swapped := false
		for j := i + 1; j < len(x); j++ {
			if x[j] < x[i] {
				x[i], x[j] = x[j], x[i]
				swapped = true
			}
		}
		// this is a nice skill, will skip all the work left if they are already sorted.
		if !swapped {
			break
		}
	}
	return
}