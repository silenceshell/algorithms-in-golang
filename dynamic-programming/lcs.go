package main

import (
	"fmt"
	"bytes"
)

// Longest Common Subsequence
func main() {
	var str1 string = "ABCDGHOOxx"
	var str2 string = "AEDFHRxOxz"

	length := lcs1(str1, str2)
	fmt.Println(length)
}

// a basic solution
func lcs1(str1, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var result bytes.Buffer

	var anchor int
	var sum int
	for i:=0; i< len1; i++ {
		for j:=anchor; j<len2; j++ {
			if str1[i] == str2[j] {
				anchor = j+1
				sum++
				result.WriteByte(str1[i])
				break
			}
		}
	}

	return result.String()
}