package main

import (
	"fmt"
	"strconv"
)

func descIter(iter []int, title string) {
	pValue := false
	if pValue {
		fmt.Printf("%s: \t tpye='%T', len='%v', cap='%v', value=%v \n", title, iter, len(iter), cap(iter), iter)
	} else {
		fmt.Printf("%s: \t tpye='%T', len='%v', cap='%v'\n", title, iter, len(iter), cap(iter))
	}

}

func main() {
	s1 := [...]int{1, 2, 3, 4, 5, 6, 7}
	s2 := s1[:]
	s3 := s2[3:5]
	s4 := []int{1, 2, 3}
	s5 := []int{}
	var s6 []int

	s3[0] = 1234

	descIter(s1[:], "s1")
	descIter(s2, "s2")
	descIter(s3, "s3")
	descIter(s4, "s4")
	descIter(s5, "s5")
	descIter(s6, "s6")
	s5 = append(s5, 7, 8)
	descIter(s5, "s5_2")

	fmt.Println("---\n\n\n")

	s10 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 10; i < 30; i++ {
		descIter(s10, "s"+strconv.Itoa(i))
		s10 = append(s10, i)
	}

}
