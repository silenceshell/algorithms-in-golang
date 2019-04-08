package main

import (
    "fmt"
)
func searchInsert_(nums []int, target int) int {
    var index int = 0
    for ; index < len(nums); index++ {
        if nums[index] >= target {
            return index
        }
    }
    return index
}
func searchInsert(nums []int, target int) int {
    l := len(nums)
    if target < nums[0] {
        return 0
    } else if target > nums[l-1] {
        return l
    }

    var gap int
    var begin = 0
    var end = l-1
    for {
        gap = (end - begin)/2
        //fmt.Println(gap)
        if gap == 0 {
            if nums[begin] == target {
                return begin
            }
            if nums[end] >= target {
                return end
            }
            break
        }
        if nums[begin+gap] == target {
            return begin+gap
        }
        if nums[begin+gap] < target {
            begin = begin+gap
        } else {
            end = begin+gap
        }
    }
    return begin+gap
}

func main() {
    fmt.Println(searchInsert([]int{1,3}, 1))
    fmt.Println(searchInsert([]int{1,3}, 3))
    fmt.Println(searchInsert([]int{1,3,5,6}, 5))
    fmt.Println(searchInsert([]int{1,3,5,6}, 2))
    fmt.Println(searchInsert([]int{1,3,5,6}, 7))
    fmt.Println(searchInsert([]int{1,3,5,6}, 0))

}
