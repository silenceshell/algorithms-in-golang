package main

import "fmt"

func main() {
	var target string = "hello this is my house and my cat, 12"
	var patterns []string = []string{"hello", "world", "123", "and", "cat"}
	fmt.Println(patterns)
	fmt.Println(target)

	rabinKarp(target, patterns)

}

func rabinKarp(target string, patterns []string) {
	n, m := len(target), minLen(patterns)

	hmap := genPatternHash(patterns, m)
	var matches map[int]int = make(map[int]int)

	var multi int = 1
	for i := 0; i < (m - 1); i++ {
		multi = multi * base
	}

	h := hash(target[:m])
	for i := 0; i < (n - m + 1); i++ {
		// h := hash(target[i:i+m])
		// use below code instead.
		if i > 0 {
			h = h - multi*int(target[i-1])
			h = h*base + int(target[i+m-1])
		}
		for _, v := range hmap[h] {
			if patterns[v] == target[i:i+len(patterns[v])] {
				matches[v] = i
			}
		}
	}

	for i, v := range matches {
		fmt.Printf("pattern %s is in position %d\r\n", patterns[i], v)
	}
}

const base = 16777619

func hash(str string) int {
	var h int
	for i := 0; i < len(str); i++ {
		h = h*base + int(str[i])
	}
	return h
}

func genPatternHash(patterns []string, minLen int) map[int][]int {
	var patternHash map[int][]int = make(map[int][]int)
	for i, v := range patterns {
		h := hash(v[:minLen])
		if _, ok := patternHash[h]; ok {
			patternHash[h] = append(patternHash[h], i)
		} else {
			patternHash[h] = []int{i}
		}
	}
	return patternHash
}

func minLen(strArrays []string) (length int) {
	length = len(strArrays[0])
	for _, v := range strArrays {
		if length > len(v) {
			length = len(v)
		}
	}
	return
}
