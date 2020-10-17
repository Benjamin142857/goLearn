package main

import (
	"fmt"

	. "goLearn/d11_package/utils"
	. "goLearn/d11_package/utils2"
)

func init() {
	fmt.Println("main.go init")
}

func main() {
	ret := Add(1, 2)
	fmt.Println(ret)

	ret2 := Plus(4, 5)
	fmt.Println(ret2)

	A()
}
