// Knuth–Morris–Pratt (KMP) in Golang
package main

import (
	"fmt"
)

const (
	PatternSize int = 100
)

// https://zh.wikipedia.org/wiki/%E5%85%8B%E5%8A%AA%E6%96%AF-%E8%8E%AB%E9%87%8C%E6%96%AF-%E6%99%AE%E6%8B%89%E7%89%B9%E7%AE%97%E6%B3%95
func KMPNew(haystack string, needle string) []int {
	var next []int = preKMPNew(needle)

	var i int = 0
	var j int = 0
	//var r int = 0
	var lenH int = len(haystack)
	var lenN int = len(needle)

	var result []int = []int{}
	var count int = 0

	for i < lenH {
		count++
		if haystack[i] != needle[j] {
			i -= next[j]
			j = 0
			//i++
			//j = next[j]
		} else {
			i++
			j++
		}

		if j >= lenN {
			result = append(result, i-lenN)
			j = 0
		}
	}
	//fmt.Println(count)

	return result
}

func preKMPNew(x string) []int {
	var i int = 1
	var j int = 0
	var length int = len(x)
	var next []int = make([]int, length)
	next[0] = -1

	for i < length {
		if x[i] == x[j] {
			next[i] = j
			j++
		} else {
			next[i] = j
			j = 0
		}
		i++
	}

	return next
}

// http://www.golangprograms.com/data-structure-and-algorithms/golang-program-for-implementation-of-knuth-morris-pratt-kmp-algorithm.html
func KMP(haystack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	count := 0
	m := len(needle)
	n := len(haystack)

	x := []byte(needle)
	y := []byte(haystack)
	var ret []int

	//got zero target or want, just return empty result
	if m == 0 || n == 0 {
		return ret
	}

	//want string bigger than target string
	if n < m {
		return ret
	}

	for j < n {
		count++
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		//fmt.Println(i, j)
		if i >= m {
			ret = append(ret, j-i)
			//fmt.Println("find:", j, i)
			i = next[i]
		}
	}
	//fmt.Println(count)

	return ret
}

func preMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var mpNext [PatternSize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}
	return mpNext
}

func preKMP(x string) [PatternSize]int {
	var i, j int
	length := len(x) - 1
	var kmpNext [PatternSize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < length {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}

		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}

func main() {
	fmt.Println("Search First Position String:")
	fmt.Println(KMP("cocacola", "ABCDABD"))
	fmt.Println(KMP("Australia", "lia"))
	fmt.Println(KMP("cocacola", "cx"))
	fmt.Println(KMP("AABAACAADAABAABA", "AABA"))

	fmt.Println("\nSearch Last Position String:\n")
	fmt.Println(KMP("cocacola", "co"))
	fmt.Println(KMP("Australia", "lia"))
	fmt.Println(KMP("cocacola", "cx"))
	fmt.Println(KMP("AABAACAADAABAABAAABAACAADAABAABA", "AABA"))

	fmt.Println("Search First Position String(new):")
	fmt.Println(KMPNew("cocacola", "ABCDABD"))
	fmt.Println(KMPNew("Australia", "lia"))
	fmt.Println(KMPNew("cocacola", "cx"))
	fmt.Println(KMPNew("AABAACAADAABAABA", "AABA"))

	fmt.Println("\nSearch Last Position String(new):")
	fmt.Println(KMPNew("cocacola", "co"))
	fmt.Println(KMPNew("Australia", "lia"))
	fmt.Println(KMPNew("cocacola", "cx"))
	fmt.Println(KMPNew("AABAACAADAABAABAAABAACAADAABAABA", "AABA"))
}
