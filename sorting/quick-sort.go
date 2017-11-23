package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/silenceshell/algorithms-in-golang/utils"
)

func main() {
	var x []int = utils.GenerateSlice(13)

	fmt.Println(x)

	fmt.Println("==========")

	quickSort(x)
	fmt.Println(x)
}

func quickSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	rand.Seed(time.Now().UnixNano())
	mid := rand.Intn(len(x))

	left, right := 0, len(x)-1
	x[mid], x[right] = x[right], x[mid]

	for i := 0; i < len(x); i++ {
		if x[i] < x[right] {
			x[left], x[i] = x[i], x[left]
			left++
		}
	}

	x[left], x[right] = x[right], x[left]

	quickSort(x[:left])
	quickSort(x[left+1:])

	return x
}
