package main

import (
	"fmt"
	"strconv"
)

func test1() {
	var mData map[string]string = make(map[string]string, 1)
	for i := 0; i < 20; i++ {
		mData[strconv.Itoa(i)] = "666"
	}
	fmt.Println(len(mData))
}

func main() {
	test1()
}
