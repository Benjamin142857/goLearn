package main

import "fmt"

type list []interface{}

func (l *list) push(elem interface{}) {
	*l = append(*l, elem)
}

func assign(a interface{}) {
	sType := fmt.Sprintf("%T", a)
	//aa, err := a.(type)
	fmt.Println(sType)
	//fmt.Println(err)
}

func assign2(a interface{}) {
	switch a.(type) {
	case list:
		fmt.Println("This is list")
	default:
		fmt.Println("Not list")
	}

}

func main() {
	l1 := list{}
	//fmt.Println(l1)
	//
	//l1.push("abc")
	//fmt.Println(l1)
	//
	//l1.push(1n  23)
	//fmt.Println(l1)
	//
	//l1.push([]int{1, 2, 3})
	//fmt.Println(l1)
	assign2(l1)
	assign2(123)
}
