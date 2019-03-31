package main

import "fmt"

func isValid(s string) bool {
    t := make([]byte, 0, len(s))
    for i:=0; i < len(s); i++ {
        //fmt.Println(s[i])
        l := len(t)
        if l != 0 && ((s[i] - t[l-1]) == 2 || ((s[i] - t[l-1]) == 1)) {
            t = t[:l-1]
        } else {
            t = append(t, s[i])
        }
        //fmt.Printf("%v\n", t)
    }
    if len(t) == 0 {
        return true
    }
    return false
}

func main() {
    fmt.Println(isValid("()"))
    fmt.Println(isValid("()[]{}"))
    fmt.Println(isValid("(]"))
    fmt.Println(isValid("(("))
    fmt.Println(isValid("{[]}"))

}
