package main

import (
	"fmt"
)

func main() {
	sTest := "Benjamin142857大帅比"
	for idx, cCode := range sTest {
		c := string(cCode)
		fmt.Println(idx, c)
	}
}
