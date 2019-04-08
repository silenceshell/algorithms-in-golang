package main

import "fmt"

func addBinary(a string, b string) string {
    var la = len(a)
    var lb = len(b)
    var l int

    if la > lb {
        l = la
    } else {
        l = lb
    }
    l+=1

    var res []byte = make([]byte, l)
    var lai byte
    var lbi byte
    var last byte = 0
    var sum byte = 0

    for i:=0; i<l; i++ {
        if la >= (i+1) {
            lai = byte(a[la-i-1]-48)
        } else {
            lai = 0
        }
        if lb >= (i+1) {
            lbi = b[lb-i-1]-48
        } else {
            lbi = 0
        }
        sum = lai + lbi +last
        if sum > 1 {
            last = 1
            res[l-i-1] = sum-2+48
        } else {
            last = 0
            res[l-i-1] = sum+48
        }
        fmt.Println(lai, lbi, last, res[l-i-1])
    }
    if res[0] == 48 {
        return string(res[1:])
    }
    return string(res)
}

func main() {
    fmt.Println(addBinary("11", "1"))
    fmt.Println(addBinary("1", "1"))
    fmt.Println(addBinary("1", "0"))
    fmt.Println(addBinary("0", "0"))
    fmt.Println(addBinary("1010", "1011"))
    fmt.Println(addBinary("1111", "1111"))

}
