package main

import "fmt"

func main() {
	a := make(map[string]int, 3)
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3
	fmt.Println(len(a))
	a["d"] = 4
	fmt.Println(a)
	fmt.Println(len(a))
}
