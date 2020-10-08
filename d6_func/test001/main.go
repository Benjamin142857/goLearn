package main

import (
	"fmt"
	"strings"
)

func test1(s []int) (ret int) {
	for _, v := range s {
		ret += v
	}
	return ret
}

func test2(x, y int, z, w string) {
	fmt.Println(x, y)
	fmt.Println(z, w)
}

func test3(x int, y ...string) {
	fmt.Println(x)
	fmt.Println(y)
}

func test4(x int, y []string) {
	fmt.Println(x)
	fmt.Println(y)
}

func test5() {
	a := [...]int{10: 200, 30: 0}
	b := [5]int{3}
	fmt.Println(len(a))
	fmt.Println(b)
}

func test6() {
	a := "abcde"
	fmt.Println(strings.Split(a, ""))
}

func judgePalindrome(s string) bool {
	s2 := strings.Join(strings.Split(s, ""), "#")
	sl := len(s2)
	for i := 0; i < (sl-1)/2; i++ {
		if s2[i] != s2[sl-1-i] {
			return false
		}
	}
	return true
}

func main() {
	//ret := test1([]int{1, 2, 3, 4})
	//fmt.Println(ret)
	//test3(10, "abc", "benjamin")
	//test3(10, []string{"abc", "benjamin"}...)
	//test4(10, []string{"abc", "benjamin"})
	ret := judgePalindrome("ABCBA")
	fmt.Println(ret)
}
