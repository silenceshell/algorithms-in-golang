package utils

import (
	"time"
	"math/rand"
)

func GenerateSlice(size int) []int {
	x := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(x); i++ {
		x[i] = rand.Intn(999) - rand.Intn(999)
	}
	return x
}
