package main

import (
	"fmt"
)

func main() {
	a := 100
	p := &a
	var p2 *int
	fmt.Println(p2)
	*p = 200

	fmt.Println(a)
}
