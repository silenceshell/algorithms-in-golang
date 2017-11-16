package main

import (
	"time"
	"math/rand"
	"fmt"
)

func main() {
	x := generateSlice(13)
	fmt.Println(x)

	bubbleSort(x)
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