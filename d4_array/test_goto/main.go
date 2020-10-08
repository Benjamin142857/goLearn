package main

import "fmt"

// 利用 goto 代替内层循环 break
func test1() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'O'; j++ {
			if j == 'C' {
				goto breakInner
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	breakInner:
	}
}

// 利用 goto 代替内层循环 continue
func test2() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'O'; j++ {
			if j == 'C' {
				goto breakInner
			}
			fmt.Printf("%v-%c\n", i, j)
		breakInner:
		}
	}
}

// 利用 goto 代替外层循环 break
func test3() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'O'; j++ {
			if j == 'C' {
				goto breakOuter
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
breakOuter:
}

// 利用 goto 代替外层循环 continue
func test4() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'O'; j++ {
			if j == 'C' {
				goto breakOuter
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	breakOuter:
	}

}

func main() {
	test4()
}
