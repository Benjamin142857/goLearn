package main

import "fmt"

type TestBody struct {
	name string
	age  int
}

func test(iTest interface{}) {
	ke, ok := iTest.(TestBody)
	if !ok {
		fmt.Println("False")
		return
	}
	fmt.Println(ke)
}

// TestBody 构造函数
func NewTestBody(options ...interface{}) *TestBody {
	pTb := new(TestBody)
	for idx, opt := range options {
		switch idx {
		case 0:
			v, ok := opt.(string)
			if !ok {
				break
			}
			pTb.name = v
		case 1:
			v, ok := opt.(int)
			if !ok {
				break
			}
			pTb.age = v
		}
	}
	return pTb
}

func main() {
	fmt.Println(NewTestBody("kkk"))
}
