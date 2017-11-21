package main

import (
	"github.com/silenceshell/algorithms-in-golang/utils"
	"fmt"
)

func main() {
	x := utils.GenerateSlice(13)
	fmt.Println(x)

	x = mergeSort(x)
	fmt.Println(x)
}

func mergeSort(x []int) []int {
	var (
		left []int
		right []int
	)
	if len(x) == 1 {
		return x
	}

	divide := len(x) / 2

	left = mergeSort(x[:divide])
	right = mergeSort(x[divide:])

	return merge(left, right)
}

func merge(a []int, b []int) []int {
	var tmp []int = []int{}

	i, j := 0, 0
	for ; i < len(a) && j < len(b); {
		if a[i] < b[j] {
			tmp = append(tmp, a[i])
			i++
		} else {
			tmp = append(tmp, b[j])
			j++
		}
	}

	tmp = append(tmp, a[i:]...)
	tmp = append(tmp, b[j:]...)

	return tmp
}