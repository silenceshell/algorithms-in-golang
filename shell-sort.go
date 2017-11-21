package main

import (
	"github.com/silenceshell/algorithms-in-golang/utils"
	"fmt"
)

func main() {
	var x []int = utils.GenerateSlice(13)
	fmt.Println(x);

	simpleShellSort(x)
	fmt.Println(x)
}

func simpleShellSort(x []int) {
	inc := len(x)/2

	for inc > 0 {
		for i:=inc; i<len(x);i++ {
			j := i
			for j > inc {
				if x[j] < x[j-inc] {
					x[j], x[j-inc] = x[j-inc], x[j]
				}
				j-=inc
			}
		}

		inc = inc/2
	}
}