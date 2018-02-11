package main

import "fmt"

func main() {
	var nums []int = []int{3, 2, 4}
	results := twoSumBest(nums, 6)
	fmt.Println(results)
}

func twoSumBest(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	var m map[int]int = make(map[int]int, len(nums))
	for i, v := range nums {
		if j, ok := m[v]; ok {
			return []int{i, j}
		} else {
			m[target-v] = i
		}
	}

	return nil
}

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	var m map[int]int = make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}

	for i, v := range nums {
		if j, ok := m[target - v]; ok {
			if i != j {
				return []int{i, j}
			}
		}
	}
	return nil
}
