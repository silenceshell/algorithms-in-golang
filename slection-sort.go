package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	x := generateSlice(13)
	fmt.Println(x)

	selectionSort(x)
	fmt.Println(x)
}

func generateSlice(size int) []int {
	x := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(999) - rand.Intn(999)
	}
	return x
}

func selectionSort(x []int) {
	for i:=0; i<len(x); i++ {
		var index int
		for j:=i+1; i< len(x); j++ {
			if x[j] < x[index] {
				index = j
			}
		}
		if x[index] < x[i] {
			x[i], x[index] = x[index], x[i]
		}
	}
}