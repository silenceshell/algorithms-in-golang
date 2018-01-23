package main

import "fmt"

//https://zh.wikipedia.org/zh-hans/%E6%B1%89%E8%AF%BA%E5%A1%94
func hanoi(h int, from string, via string, to string) {
	if h == 1 {
		fmt.Printf("move %d from %s to %s\n", h, from, to)
		return
	}
	hanoi(h-1, from, to, via)
	fmt.Printf("move %d from %s to %s\n", h, from, to)
	hanoi(h-1, via, from, to)
}

func main() {
	hanoi(3, "A", "B", "C")
}
