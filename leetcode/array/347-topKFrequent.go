package main

import "fmt"

func main() {
	var x []int = []int{1, 1, 1, 2, 2, 3}
	fmt.Println(topKFrequent(x, 2))

	x = []int{3, 0, 1, 0}
	fmt.Println(topKFrequent(x, 1))

}
type resA struct {
	value int
	count int
}
func topKFrequent(nums []int, k int) []int {
	var x map[int]int = make(map[int]int)
	for _, n := range nums {
		if v, ok := x[n]; ok {
			x[n] = v+1
		} else {
			x[n] = 1
		}
	}

	var res []resA = make([]resA, k)
	for k, v := range x {
		if v > res[0].count {
			insertToRes(k, v, res)
		}
	}

	var r[]int = make([]int, k)
	for i, v := range res {
		r[i] = v.value
	}

	return r
}
func insertToRes(k, v int, res []resA) {
	var i int = 1
	for ; i < len(res); i++ {
		if v > res[i].count {
			res[i-1] = res[i]
		} else {
			res[i-1].value = k
			res[i-1].count = v
			break
		}
	}
	if i==len(res) {
		res[i-1].value = k
		res[i-1].count = v
	}
}