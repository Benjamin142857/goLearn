package main

import (
	"fmt"
)

func main() {
	a := 100
	b := uint8(200)
	c := 0x123af
	d := 0o123

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T \n", b)

	fmt.Println(c)
	fmt.Println("========================")
	fmt.Printf("%b \n", a)
	fmt.Printf("%o \n", a)
	fmt.Printf("%d \n", a)
	fmt.Printf("%x \n", a)

	fmt.Println("========================")
	fmt.Println(d)
}
