package main

import (
	"fmt"
)

func changeS0(s []int, x int) {
	s[0] = x
}

func test1() (ret []int) {
	fmt.Println("start")
	ret = []int{99}
	for i := 0; i < 20; i++ {
		defer changeS0(ret, i)
	}
	fmt.Println("end")
	return
}

func f1() int {
	x := 5
	defer func() { x++ }()
	return x
}

func f2() (x int) {
	x = 0
	defer func() { x++ }()
	return 5
}

func f3() (x int) {
	x = 5
	defer func(x int) { x++ }(x)
	return
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
