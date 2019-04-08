package main

import (
	"fmt"
)

/*
Given a string, find the length of the longest substring without repeating characters.

Examples:

Given "abcabcbb", the answer is "abc", which the length is 3.
Given "bbbbb", the answer is "b", with the length of 1.
Given "pwwkew", the answer is "wke", with the length of 3. Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/

func main() {
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdf"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("ckilbkd"))
	fmt.Println(lengthOfLongestSubstring("aab"))
}

func lengthOfLongestSubstring(s string) int {
	var x map[int32]int = make(map[int32]int, 256)
	var len int = 0
	var max int = 0
	var beginIndex int = 0
	for index, value := range s {
		if valueIndex, ok := x[value]; !ok {
			x[value] = index
			len++
		} else {
			if len > max {
				max = len
			}

			for j:=beginIndex; j <= valueIndex; j++ {
				delete(x, int32(s[j]))
			}
			len -= (valueIndex - beginIndex)
			beginIndex = valueIndex +1
		}
	}
	if len > max {
		return len
	} else {
		return max
	}
}
