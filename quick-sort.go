package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	var x []int = generateSlice(13)

	for _, v := range x {
		fmt.Println(v)
	}

	fmt.Println("==========")
	quickSort(x)

	for _, v := range x {
		fmt.Println(v)
	}
}

func generateSlice(size int) []int {
	x := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i:=0; i< len(x); i++ {
		x[i] = rand.Intn(999) - rand.Intn(999)
	}
	return x
}

func quickSort(x []int) []int {
	if len(x) < 2 {
		return x
	}
	rand.Seed(time.Now().UnixNano())
	mid := rand.Intn(len(x))

	left, right := 0, len(x) -1
	x[mid], x[right] = x[right], x[mid]

	for i:=0; i< mid; i++ {
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