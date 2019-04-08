package main

import "fmt"

func romanToInt(s string) int {
	var ret,last,num int
	for _, v := range s {
		switch string(v) {
		case "I":
			num = 1
		case "V":
			num = 5
		case "X":
			num = 10
		case "L":
			num = 50
		case "C":
			num = 100
		case "D":
			num = 500
		case "M":
			num = 1000
		}
        ret += num
        if last < num {
            ret -= last*2
        }
        last = num
	}
	return ret
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
