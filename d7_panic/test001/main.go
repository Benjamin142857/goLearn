package main

import "fmt"

func fnA() {
	fmt.Println("A")
}

func fnB() {
	//defer func() {
	//	err := recover()
	//	fmt.Printf("%T\n", err)
	//	fmt.Println(err)
	//}()
	//a := []int{1, 2}
	//fmt.Println(a[10])
	//fmt.Println("B")
	panic("出错啦！！！")
}

func fnC() {
	fmt.Println("C")
}

func main() {
	fnA()
	fnB()
	fnC()
}
