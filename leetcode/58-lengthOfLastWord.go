package main

import "fmt"

func lengthOfLastWordDP(s string) int {
    var l int = 0
    var ls = len(s)
    if ls == 0 {
        return 0
    }
    var c []int = make([]int, ls)
    if s[0] == 32 {
        c[0] = 0
    } else {
        c[0] = 1
    }
    for i:=1; i<ls; i++ {
        if s[i] != 32 {
            l = c[i-1] + 1
        } else {
            l = 0
        }
        c[i] = l
    }

    for i:= ls-1; i>=0; i-- {
        if c[i] != 0 {
            return c[i]
        }
    }
    return 0
}

func lengthOfLastWord(s string) int {
    var ls = len(s)
    var cnt = 0
    for i:=ls-1; i>=0; i-- {
        if s[i] != 32 {
            cnt++
        } else if cnt != 0 {
            return cnt
        }
    }
    return cnt
}

func main() {
    fmt.Println(lengthOfLastWord("Hello World"))
    fmt.Println(lengthOfLastWord("Hello World "))
    fmt.Println(lengthOfLastWord("Today is a nice day"))
    fmt.Println(lengthOfLastWord("b   a    "))
    fmt.Println(lengthOfLastWord(""))
    fmt.Println(lengthOfLastWord("a"))
    fmt.Println(lengthOfLastWord(" "))
}
