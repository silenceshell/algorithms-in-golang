package main

import (
	"sort"
	"fmt"
)

func main() {
	intSlice := []int{5, 9, 77, 62, 71, 11, 22, 46, 36, 18, 19, 33, 75, 17, 39, 41, 73, 50, 217, 79, 120}
    sort.Ints(intSlice)
	fmt.Println(intSlice)
}