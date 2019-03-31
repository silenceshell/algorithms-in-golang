package main

import "fmt"

func strStr(haystack string, needle string) int {
    if needle == "" {
        return 0
    }
    l := len(haystack) - len(needle) + 1
	for i := 0; i < l; i++ {
        var j int
		for j=0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
        if j == len(needle) {
            return i
        }
	}
	return -1
}
func main() {
	fmt.Println(strStr("aaa", "aaaa"))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("aaaaa", "aab"))
	fmt.Println(strStr("", "aab"))
	fmt.Println(strStr("iiii", ""))
	fmt.Println(strStr("", ""))

}
