package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	var items []int = generateSlice(13)
	for i := range items {
		fmt.Printf("%d ", items[i])
	}
	fmt.Println()
	var target = items[rand.Intn(len(items))]
	fmt.Printf("searching target is: %d\r\n", target)
	index, err := linearSearch(items, target)
	if err == nil {
		fmt.Printf("found %d, index is %d", target, index)
	} else {
		fmt.Printf("%d not found", target)
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

func linearSearch(items []int, target int) (int, error) {
	//var index int
	for index := range(items) {
		if items[index] == target {
			return index, nil
		}
	}
	return 0, fmt.Errorf("%d Not found", target)
}