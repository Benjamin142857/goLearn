package main

import (
	"fmt"
	"math"
)

func main() {
	a := 12.3
	b := 4.56
	c := 1 + 3i

	fmt.Printf("%v\n", math.MaxFloat32)
	a = b
	fmt.Println(a)
	fmt.Println(c * c)
	fmt.Printf("%v\n", c)
	q := true

	fmt.Printf("%T, %v \n", q, q)
}
