package main

import "fmt"

const a = iota
const b = iota

const (
	q1 = iota
	q2 = iota
	_  = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(q1)
	fmt.Println(q2)
}
