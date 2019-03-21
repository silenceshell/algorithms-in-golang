package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
    if len(strs) == 1 {
        return strs[0]
    }
	var index int
	for {
		for count := 0; count < len(strs)-1; count++ {
			if index >= len(strs[count]) || index >= len(strs[count+1]) || strs[count][index] != strs[count+1][index] {
				return strs[count][:index]
			}
			//            if strs[count][index] != strs[count+1][index] {
			//                return strs[count][:index]
			//            }
		}
		index += 1
	}
	return ""
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{}))
	fmt.Println(longestCommonPrefix([]string{""}))
	fmt.Println(longestCommonPrefix([]string{"", ""}))
	fmt.Println(longestCommonPrefix([]string{"aa", "a"}))
}
