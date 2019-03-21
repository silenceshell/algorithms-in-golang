package main

import "fmt"

func reverse(x int32) int32 {
    var pop int32
    var y int32 = 0
    for {
        pop = x%10
        x = x/10
        if (y > 0x7fffffff/10) || (y < -0x7fffffff/10) || (y==0x7fffffff/10 && pop > 7) || (y==-0x7fffffff/10 && pop < -8) {
            return 0
        }
        y = (y*10 + pop)
        if x == 0 {
            break
        }
    }

    return y
}

func main() {
    re := reverse(123)
    if re != 321 {
        fmt.Printf("err %v\n", re)
    }
    re = reverse(1)
    if re != 1 {
        fmt.Printf("err %v\n", re)
    }
    re = reverse(1534236469)
    if re != 0 {
        fmt.Printf("err %v\n", re)
    }
    re2 := reverse(-123)
    if re2 != -321 {
        fmt.Printf("err %v\n", re2)
    }
}
