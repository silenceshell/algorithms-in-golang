package main

import "fmt"

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    var origin int = x
    var pop int
    var y int = 0
    for {
        pop = x%10
        x = x/10
        if (y > 0x7fffffff/10) || (y < -0x7fffffff/10) || (y==0x7fffffff/10 && pop > 7) || (y==-0x7fffffff/10 && pop < -8) {
            return false
        }
        y = (y*10 + pop)
        if x == 0 {
            break
        }
    }

    return y==origin
}

func main() {
    fmt.Println(isPalindrome(121))
    fmt.Println(isPalindrome(-121))
    fmt.Println(isPalindrome(10))
}
