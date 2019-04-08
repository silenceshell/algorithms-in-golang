package main

import "fmt"

func maxSubArray(nums []int) int {
	sums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sums[i] = -1
	}
    sums[0] = nums[0]
    var max int = nums[0]

	for i := 1; i < len(nums); i++ {
        if sums[i-1] > 0 {
            sums[i] = sums[i-1] + nums[i]
        } else {
            sums[i] = nums[i]
        }
        if max < sums[i] {
            max = sums[i]
        }
	}

    return max
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(arr))
}
