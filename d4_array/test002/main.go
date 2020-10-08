package main

import (
	"fmt"
)

func switchDemo(key int) {
	switch key {
	case 1:
		fmt.Println("Key=1")
	case 2:
		fmt.Println("Key=2")
	case 3:
		fmt.Println("Key=3")
		fallthrough
	case 4:
		fmt.Println("Key=4")
		fallthrough
	default:
		fmt.Println("DEFAULT")
	}
}

func main() {
	key := 3
	switchDemo(key)
}
