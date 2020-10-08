package main

import (
	"fmt"
)

func main() {
	strTest := "我儿子古富源是一个大傻子"
	for _, v := range strTest {
		if v == '子' {
			fmt.Println(string(v))
		}
	}
}
