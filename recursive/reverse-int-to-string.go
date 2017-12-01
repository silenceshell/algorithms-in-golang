package main

import (
	"math/rand"
	"time"
	"fmt"
)

// reverse an int, like 12345, to string "54321"
func main() {
	rand.Seed(time.Now().UnixNano())
	var x int = rand.Int()
	fmt.Println(x)

	var result string = reverseIntToString(x)
	fmt.Println(result)
}

func reverseIntToString(data int) string {
	var s string = string((data % 10) + '0')
	if data > 10 {
		return s + reverseIntToString(data/10)
	} else {
		return s
	}
}
