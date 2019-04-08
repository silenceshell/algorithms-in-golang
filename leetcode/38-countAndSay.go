package main

import (
    "fmt"
)

func convert(s string) string {
    var c int = 1
    var ne string = ""
    if len(s) == 1 {
        return "1"+s
    }
    for i := 0; i<len(s)-1; i++ {
        if s[i] != s[i+1] {
            ne = ne + string(48+c) + string(s[i])
            c = 1
        } else {
            c++
        }
    }
    ne = ne + string(48+c) + string(s[len(s)-1])
    return ne
}

func countAndSay(n int) string {
    if n == 1 {
        return "1"
    } else {
        return convert(countAndSay(n-1))
    }
}

func main() {
    fmt.Println(countAndSay(3))
    fmt.Println(countAndSay(4))
    fmt.Println(countAndSay(5))
}
