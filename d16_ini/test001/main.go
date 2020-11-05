package main

import (
	"fmt"
	"strings"
)

func test1() {
	testMsg := "   \tBenjamin142857   \t   \n  "

	fmt.Println(strings.TrimSpace(testMsg))
}
func main() {
	test1()
}
