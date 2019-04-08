package main

import "fmt"

func lengthOfLastWord(s string) int {
    var l int = 0
    var last int = 0
    for _, v := range s {
        //fmt.Println(k, v)
        if v != 32 {
            l++
        } else {
            if l != 0 {
                last = l
                l = 0
            }
        }
    }
    if last != 0 && last != l {
        last = l
    }
    return last
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
