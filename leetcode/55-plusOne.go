package main

import "fmt"
func plusOne(digits []int) []int {
    var ls = len(digits)
    var res []int = make([]int, ls+1)
    var last = 1
    var sum = 0
    for i:= ls-1; i>=0; i-- {
        sum = last + digits[i]
        if sum >= 10 {
            last = 1
            sum = sum-10
        } else {
            last = 0
        }
        res[i+1] = sum
    }
    if last == 1 {
        res[0] = 1
        return res
    } else {
        return res[1:]
    }
}

func main() {
    fmt.Println(plusOne([]int{1,2,3}))
    fmt.Println(plusOne([]int{4,3,2,1}))
    fmt.Println(plusOne([]int{1}))
    fmt.Println(plusOne([]int{9,9}))
    fmt.Println(plusOne([]int{1,9}))
}
