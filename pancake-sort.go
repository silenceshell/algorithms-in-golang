package main

import (
	"github.com/silenceshell/algorithms-in-golang/utils"
	"fmt"
)

func main() {
	x := utils.GenerateSliceInt32(13)
	fmt.Println(x)

	pancakeSort(x)
	fmt.Println(x)
}

func pancakeSort(x []int32) {
	for length := len(x); length > 0; length-- {
		// find the biggest pancake
		maxIndex, max := 0, x[0]
		for i:=0; i < length; i++ {
			if x[i] > max {
				maxIndex = i
				max = x[i]
			}
		}

		// flip twice
		flip(x, maxIndex)
		flip(x, length-1)
	}
}

func flip(x []int32, n int) {
	for left, right:=0, n; left < right; left, right=left+1, right-1 {
		x[left], x[right] = x[right], x[left]
	}
}