package main

import "fmt"

// Animal 一个动物单元
type Animal struct {
	name   string
	class  string
	age    int
	volume float64
	hobby  []string
}

type Book struct {
	name   string
	price  int
	author string
}

func test1() {
	a := Animal{"Alex", "dog", 20, 100, []string{"play", "eat"}}
	b := Animal{
		name:   "Benjamin",
		class:  "cat",
		age:    20,
		volume: 155,
		hobby:  []string{"play", "study"},
	}
	c := new(Animal)
	c.name = "stella"
	c.class = "pig"
	c.age = 22
	c.volume = 1000
	c.hobby = []string{"eat", "eat", "eat"}
	fmt.Printf("%T, %+v\n", a, a)
	fmt.Printf("%T, %+v\n", b, b)
	fmt.Printf("%T, %+v\n", *c, *c)

}

func test2() {
	a := 100
	fmt.Printf("%v", &a)
}

func newAnimal(name string) Animal {
	return Animal{
		name:   name,
		class:  "human",
		age:    18,
		volume: 100,
		hobby:  make([]string, 10),
	}
}

func (A Animal) speak() {
	fmt.Printf("%s 大叫了一声...\n", A.name)
	A.volume--
}
func (A *Animal) eat(food string) {
	fmt.Printf("%s 吃了%s...\n", A.name, food)
	A.volume++
}

func test3() {
	a := newAnimal("Stella")
	b := newAnimal("Benjamin")
	fmt.Printf("Type(a)=%T\n", a)
	fmt.Printf("Type(b)=%T\n", b)

	fmt.Println(a.volume, b.volume)
	a.speak()
	b.speak()
	fmt.Println(a.volume, b.volume)

	fmt.Println(a.volume, b.volume)
	a.eat("肥猪肉")
	b.eat("芝底榴莲薄脆披萨")
	fmt.Println(a.volume, b.volume)
}

func main() {
	test3()
}
