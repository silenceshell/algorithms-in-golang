package utils

import (
	"math/rand"
	"time"
)

func GenerateSlice(size int) []int {
	x := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(999) - rand.Intn(999)
	}
	return x
}

func GenerateSliceInt32(size int) []int32 {
	x := make([]int32, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Int31n(999) - rand.Int31n(999)
	}
	return x
}
