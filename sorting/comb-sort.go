package main

import (
	"github.com/silenceshell/algorithms-in-golang/utils"
	"fmt"
)

func main() {
	var x []int = utils.GenerateSlice(13)
	fmt.Println(x)

	combSort(x)
	fmt.Println(x)
}

func combSort(x []int) {
	var (
		length int = len(x)
	)
	for comb := length; comb >= 1; comb-- {
		for i := 0; i < length - comb; i++ {
			if x[i] > x[i+comb] {
				x[i], x[i+comb] = x[i+comb], x[i]
			}
		}
	}
}