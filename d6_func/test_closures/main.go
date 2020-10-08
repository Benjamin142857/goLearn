package main

import "fmt"

func test1(x int, y ...int) int {
	return 10
}

func test2(f func(int, ...int) int) {
	fmt.Println(f(10, 20, 30))
}

func f1(x, y int) int {
	fmt.Println("do something")
	return x + y
}
func f2(fn func(int, int) int) func(int, int) int {
	return func(x, y int) int {
		fmt.Println("start")
		ret := fn(x, y)
		fmt.Println("end")
		return ret
	}
}

func add(init int) func(int) int {
	cnt := init
	return func(x int) int {
		cnt += x
		return cnt
	}
}

func f1(x int) int {
	return 30
}

func main() {
	//f1 := f2(f1)
	//fmt.Println(f1(3, 5))

	a1 := add(100)
	fmt.Println(a1(10))
	fmt.Println(a1(20))
	fmt.Println(a1(50))

	a2 := add(0)
	fmt.Println(a2(20))
	fmt.Println(a2(80))
}
