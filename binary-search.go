package main

import "fmt"

func main() {
	var x []int = []int{1, 3, 22, 32, 40, 123, 344, 555}
	for _, v := range x {
		fmt.Printf("%d ", v)
	}
}
