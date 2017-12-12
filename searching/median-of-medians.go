package main

import (
	"fmt"
	"sort"
)

func main() {
	intSlice := []int{5, 9, 77, 62, 71, 11, 22, 46, 36, 18, 19, 33, 75, 17, 39, 41, 73, 50, 217, 79, 120}

	var newSlice []int = make([]int, len(intSlice))
	copy(newSlice, intSlice)
	sort.Ints(newSlice)
	fmt.Println(newSlice)

	for _, j := range []int{5, 10, 15, 20} {
		// search the jth smallest item in intSlice.
		i := medianOfMedians(intSlice, j, 5)
		fmt.Println(j, "smallest element = ", i)
		v := newSlice[j-1]
		fmt.Println("arr[", j-1, "] = ", v)
		if i != v {
			fmt.Println("Oops! Algorithm is wrong")
		}
	}
}

func medianOfMedians(sliceList []int, k, r int) int {
	// use median of medians to get pivot for this slice List
	length := len(sliceList)
	if length < 2*r {
		sort.Ints(sliceList)
		return sliceList[k-1]
	}
	medianLen := (length + r) / r
	var medianSlice []int = make([]int, medianLen)
	var sliceR []int
	for i := 0; i < medianLen; i++ {
		if i < (medianLen - 1) {
			sliceR = make([]int, r)
			copy(sliceR, sliceList[i*r:(i+1)*r])
		} else {
			sliceR = make([]int, length-i*r)
			copy(sliceR, sliceList[i*r:])
		}
		sort.Ints(sliceR)
		medianSlice[i] = sliceR[len(sliceR)/2]
	}
	pivot := medianOfMedians(medianSlice, (medianLen+1)/2, r)

	//anchor always point to the item >= pivot
	anchor := 0
	for i := 0; i < length; i++ {
		if sliceList[i] < pivot {
			sliceList[i], sliceList[anchor] = sliceList[anchor], sliceList[i]
			anchor++
		}
	}

	if k-1 == anchor {
		return sliceList[anchor]
	} else if k-1 < anchor {
		return medianOfMedians(sliceList[0:anchor], k, r)
	} else {
		return medianOfMedians(sliceList[anchor+1:], k-anchor-1, r)
	}

}
