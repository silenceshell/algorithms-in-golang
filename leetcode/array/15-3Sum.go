package main

import "fmt"

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {


	var ret [][]int = make([][]int, 1)
	for i, v := range nums {
		result := twoSum(append(nums[:i], nums[i+1:]...), v)
		if result != nil {
			ret[i] = append(result, i)
			return ret
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
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
