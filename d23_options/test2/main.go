package main

import (
	"fmt"
	"sort"
)

// 1. 定义结构体类型
type Person struct {
	name  string
	age   int
	hobby []string
}

// 2. 声明选项函数 (仅声明函数类型, 在参数赋值函数中被定义)
type Options func(p *Person)

// 3. 定义构造函数
func NewPerson(opts ...Options) *Person {
	p := &Person{
		name:  "none",
		age:   18,
		hobby: make([]string, 5),
	}
	for _, opts := range opts {
		opts(p)
	}
	return p
}

// 4. 定义参数赋值函数
func withName(name string) Options {
	return func(p *Person) {
		p.name = name
	}
}
func withAge(age int) Options {
	return func(p *Person) {
		p.age = age
	}
}
func withHobby(hobby []string) Options {
	return func(p *Person) {
		p.hobby = hobby
	}
}

func main() {
	p1 := NewPerson()
	fmt.Println(*p1)

	p2 := NewPerson(withName("Benjamin"), withHobby([]string{"football", "ml"}))
	fmt.Println(*p2)

	sort.IntsAreSorted()
}
