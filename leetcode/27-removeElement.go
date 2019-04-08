package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	result := removeElement(nums, 3)
	fmt.Println(nums)
	fmt.Println(result)
}

func removeElement(nums []int, val int) int {
	n := 0
	l := len(nums)
	for i:=0; i < l; i++ {
		if nums[i] != val {
			nums[n] = nums[i]
			n++
		}
	}
	return n
}