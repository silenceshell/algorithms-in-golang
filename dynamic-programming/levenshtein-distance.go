package main

import "fmt"

func main() {
	var x1 string = "Asheville"
	var x2 string = "Arizona"
	distance := levenshtein(x1, x2)
	fmt.Printf("distance between %s and %s is: %d\r\n", x1, x2, distance)
}

// this algorithm is from https://zh.wikipedia.org/wiki/%E8%90%8A%E6%96%87%E6%96%AF%E5%9D%A6%E8%B7%9D%E9%9B%A2
func levenshtein(str1, str2 string) int {
	len1 := len(str1)
	len2 := len(str2)

	var d [][]int = make([][]int, len1)
	for i := range d {
		d[i] = make([]int, len2)
	}

	for i := 0; i < len1; i++ {
		d[i][0] = i
	}
	for j := 0; j < len2; j++ {
		d[0][j] = j
	}

	var cost int
	for i := 1; i < len1; i++ {
		for j := 1; j < len2; j++ {
			if str1[i] == str2[j] {
				cost = 0
			} else {
				cost = 1
			}

			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+cost)
		}
	}

	return d[len1-1][len2-1]
}

func min(a, b, c int) int {
	if a < b && a < c {
		return a
	} else if b < a && b < c {
		return b
	} else {
		return c
	}
}
