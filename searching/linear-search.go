package main

import (
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
	"math/rand"
)

func main() {
	var items []int = utils.GenerateSlice(13)
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

func linearSearch(items []int, target int) (int, error) {
	//var index int
	for index := range items {
		if items[index] == target {
			return index, nil
		}
	}
	return 0, fmt.Errorf("%d Not found", target)
}
