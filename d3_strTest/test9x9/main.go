package main

import (
	"fmt"
)

func print9x9() {
	for i := 1; i < 10; i++ {
		for j := 1; j < i+1; j++ {
			fmt.Printf("%dx%d=%d\t", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}

func main() {
	print9x9()
}
