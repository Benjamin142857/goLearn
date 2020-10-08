package main

import "fmt"

type Animal struct {
	name   string
	class  string
	age    int
	volume float64
	hobby  []string
}

func test1() {
	a := Animal{"Alex", "dog", 20, 100, []string{"play", "eat"}}
	fmt.Printf("%T, %+v", a, a)
}

func main() {
	test1()
}
