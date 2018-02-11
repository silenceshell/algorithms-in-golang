package main

import (
	"fmt"
)

func main() {
	var nums []int = []int{1, 1, 2, 3, 3, 4}
	fmt.Println(nums)
	var newLen int = removeDuplicates(nums)
	fmt.Println(nums)
	fmt.Println(newLen)
}

func removeDuplicates(nums []int) int {
	l := len(nums)
	if l < 1 {
		return l
	}
	var index int = 0
	for i := 1; i < l; i++ {
		if nums[index] != nums[i] {
			index++
			nums[index] = nums[i]
		}
	}
	return index+1
}

func removeDuplicatesForNoSortedArray(nums []int) int {
	var index int = 0
	for i := 0; i < len(nums); i++ {
		var find bool = false
		for j := 0; j < i; j++ {
			if nums[i] == nums[j] {
				find = true
				break
			}
		}
		if !find {
			nums[index] = nums[i]
			index++
		}
	}
	return index
}

func removeDuplicatesWithHash(nums []int) int {
	var m map[int]int = make(map[int]int)
	for _, v := range nums {
		m[v] = 1
	}

	var i int = 0
	for k, _ := range m {
		nums[i] = k
		i++
	}
	return i
}
