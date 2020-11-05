package main

import (
	"fmt"
	"strconv"
)

func test1() {
	a := 69
	b := string(a)
	fmt.Println(a)
	fmt.Println(b)
}

func test2() {
	s1 := "12000"
	s2 := "133.6546"

	r1, _ := strconv.ParseInt(s1, 10, 64)
	r2, _ := strconv.ParseInt(s1, 3, 64)
	r3, err := strconv.ParseInt(s2, 10, 64)
	fmt.Printf("%T, %v\n", r1, r1)
	fmt.Printf("%T, %v\n", r2, r2)
	fmt.Printf("%T, %v\n", r3, r3)
	fmt.Println(err)
}

func main() {
	test2()
}
