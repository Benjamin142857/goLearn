package main

import "fmt"

func test1(x int) int {
	return x + 10
}

func test2(x int) (y int) {
	y = x * 2
	return
}

func main() {
	a := test1
	b := test2
	fmt.Printf("type(a)=%T\n", test1)
	fmt.Printf("type(b)=%T\n", b)
	c1 := fmt.Sprintf("%T", a)
	c2 := fmt.Sprintf("%T", b)
	fmt.Println(c1 == c2)

	fmt.Println(a(100))
	fmt.Println(b(100))
}
