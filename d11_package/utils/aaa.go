package utils

import (
	"fmt"
	_ "goLearn/d11_package/utils2"
)

func init() {
	fmt.Println("aaa.go")
}

func Plus(x, y int) int {
	return x * y
}
