package main

import "fmt"

// 阶乘
func f(n int) int {
	if n <= 1 {
		return 1
	}
	return n * f(n-1)
}

// 上台阶
func f2(n uint32) uint32 {
	if n <= 1 {
		return 1
	}
	return f2(n-1) + f2(n-2)
}

func main() {
	var a rune
	fmt.Println(f2(4))
}
