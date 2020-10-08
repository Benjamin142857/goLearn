package main

import (
	"fmt"
)

func descIter(iter []int, title string) {
	fmt.Printf("%s: \t tpye='%T', len='%v', cap='%v', value=%v \n", title, iter, len(iter), cap(iter), iter)
}

func main() {
	s1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("\n>>>>>>>>>>")
	fmt.Println(s1)
	s2 := s1[2:5]
	s3 := s2[:2]
	descIter(s2, "s2")
	descIter(s3, "s3")

	s3 = append(s3, 666)
	//s2 = append(s2, 10, 11, 13)

	descIter(s2, "s2")
	descIter(s3, "s3")

	s4 := make([]int, 7)
	descIter(s4, "s4")
	fmt.Println(">>>>>>>>>>")

}
