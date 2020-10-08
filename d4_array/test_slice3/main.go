package main

import (
	"fmt"
)

func descIter(iter []int, title string) {
	fmt.Printf("[%p] %s: \t len=%+v, cap=%+v,   %v\n", iter, title, len(iter), cap(iter), iter)
}

func test1() {
	s1 := make([]int, 5, 10)
	s2 := []int{1, 2, 3}
	s3 := [...]int{1, 2, 3}
	s4 := s3[:2]

	descIter(s1, "s1")
	descIter(s2, "s2")
	descIter(s4, "s4")
	s4 = append(s4, 123, 234)
	descIter(s4, "s4")

	var s5 []int
	if s4[3] == 0 {
		s5 = s4
	}
	fmt.Println(s5 == nil)
}

func test2() {
	s1 := make([]int, 5, 5)
	s2 := append(s1, 12, 23, 34, 45, 56, 67, 78, 89)
	descIter(s1, "s1")
	descIter(s2, "s2")
	s1 = append(s1, s2...)
	descIter(s1, "s1")
	descIter(s2, "s2")

}

func test3() {
	s1 := []int{1, 2, 3, 4, 5, 6, 7}
	s2 := make([]int, 7, 10)
	copy(s2, s1)
	descIter(s1, "s1")
	descIter(s2, "s2")

	s3 := append(s2[:3], s2[5:]...)
	descIter(s2, "s2")
	descIter(s3, "s3")

	fmt.Printf("%p\n", s3)
	fmt.Printf("%p\n", &s3[0])
}

func test4() {
	a := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	descIter(a, "a")
}

func main() {
	test4()
}
