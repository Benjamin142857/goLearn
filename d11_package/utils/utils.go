package utils

import "fmt"

func init() {
	fmt.Println("utils.go init")
}

func Add(x, y int) int {
	return x + y
}

func hello() {
	fmt.Println("Hello")
}
