package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	b := a[:]
	b[3] = 6
	fmt.Println(a)
}