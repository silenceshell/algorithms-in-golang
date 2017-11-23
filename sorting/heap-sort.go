package main

import (
	"github.com/silenceshell/algorithms-in-golang/utils"
	"fmt"
)

func main() {
	x := utils.GenerateSlice(13)
	fmt.Println(x)

	heapSort(x)
	fmt.Println(x)
}

func heapSort(x []int) {
	length := len(x)

	for m := length/2; m >=0; m-- {
		heap(x, m, length - 1)
	}
	fmt.Println(x)

	for i:=length - 1; i > 0; i-- {
		x[i], x[0] = x[0], x[i]
		heap(x, 0, i-1)
	}
}

func heap(x []int, root, end int) {
	left := 2 * root + 1
	if left > end {
		return
	}

	right := 2 * root + 2
	if right < end && x[left] < x[right] {
		left = right
	}

	if x[root] > x[left] {
		return
	}

	x[root], x[left] = x[left], x[root]
	heap(x, left, end)

}