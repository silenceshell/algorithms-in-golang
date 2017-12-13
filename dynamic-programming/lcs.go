package main

import (
	"bytes"
	"fmt"
)

// Longest Common Subsequence
func main() {
	var str1 string = "ABCDGHOOxx"
	var str2 string = "AEDFHRxOxz"

	length := lcs1(str1, str2)
	fmt.Println(length)

	l, s := lcs2(str1, str2)
	fmt.Println(l, s)
}

// a basic solution
func lcs1(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var result bytes.Buffer

	var anchor int
	var sum int
	for i := 0; i < len1; i++ {
		for j := anchor; j < len2; j++ {
			if str1[i] == str2[j] {
				anchor = j + 1
				sum++
				result.WriteByte(str1[i])
				break
			}
		}
	}

	return result.String()
}

func lcs2(str1, str2 string) (string, int) {
	len1 := len(str1)
	len2 := len(str2)
	var result bytes.Buffer

	var len [][]int = make([][]int, len1)
	for i := range len {
		len[i] = make([]int, len2)
	}

	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if i == 0 || j == 0 {
				len[i][j] = 0
			} else if str1[i] == str2[j] {
				result.WriteByte(str1[i])
				len[i][j] = len[i-1][j-1] + 1
			} else {
				len[i][j] = max(len[i-1][j], len[i][j-1])
			}
		}
	}

	return result.String(), len[len1-1][len2-1]
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
