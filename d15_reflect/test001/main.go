package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string `ini:"name"`
	Age  int    `ini:"age"`
}

type mapList []map[string]int

func test1() {
	a := 142857
	b := "Benjamin"
	c := mapList{
		{"c1": 111},
		{"c2": 222},
	}
	d := Test{"Benjamin", 22}

	va := reflect.TypeOf(a)
	vb := reflect.TypeOf(b)
	vc := reflect.TypeOf(c)
	vd := reflect.TypeOf(d)

	fmt.Printf("%T, %v\n", va, va)
	fmt.Printf("%T, %v\n", vb, vb)
	fmt.Printf("%T, %v\n", vc, vc)
	fmt.Printf("%T, %v\n", vd, vd)
}

func test2() {
	a := 142857
	b := "Benjamin"
	c := mapList{
		{"c1": 111},
		{"c2": 222},
	}

	va := reflect.TypeOf(a)
	vb := reflect.TypeOf(b)
	vc := reflect.TypeOf(c)

	fmt.Println(va.Size(), va.String())
	fmt.Println(vb.Size(), vb.String())
	fmt.Println(vc.Size(), vc.String())
}

func main() {
	test1()
}
