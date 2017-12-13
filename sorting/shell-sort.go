package main

import (
	"fmt"
	"github.com/silenceshell/algorithms-in-golang/utils"
)

func main() {
	var x []int = utils.GenerateSlice(13)
	fmt.Println(x)

	simpleShellSort(x)
	fmt.Println(x)

	var y []int = utils.GenerateSlice(13)
	fmt.Println(y)

	advanceShellSort(y)
	fmt.Println(y)
}

func simpleShellSort(x []int) {
	inc := len(x) / 2

	for inc > 0 {
		for i := inc; i < len(x); i++ {
			j := i
			for j > inc {
				if x[j] < x[j-inc] {
					x[j], x[j-inc] = x[j-inc], x[j]
				}
				j -= inc
			}
		}

		inc = inc / 2
	}
}

func square(base, square int) int {
	result := 1
	for ; square > 0; square-- {
		result = result * base
	}
	return result
}

func advanceShellSort(x []int) {
	// https://zh.wikipedia.org/wiki/%E5%B8%8C%E5%B0%94%E6%8E%92%E5%BA%8F
	// (1, 5, 19, 41, 109,...)
	// from 9*4^i - 9*2^i + 1  and 2^(i+2) * (2^(i+2) - 3) + 1
	// this is ugly.
	gaps := []int{}
	gap := 0
	for i := 0; gap < len(x); {
		gap = (9 * square(4, i)) - 9*square(2, i) + 1
		gaps = append([]int{gap}, gaps...)
		gap = square(2, i+2)*(square(2, i+2)-3) + 1
		gaps = append([]int{gap}, gaps...)
		i++
	}

	for index := 0; index <= len(gaps)-1; index++ {
		inc := gaps[index]
		for i := inc; i < len(x); i++ {
			j := i
			for j >= inc {
				if x[j] < x[j-inc] {
					x[j], x[j-inc] = x[j-inc], x[j]
				}
				j -= inc
			}
		}
	}

}
