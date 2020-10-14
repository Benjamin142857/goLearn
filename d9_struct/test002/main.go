package main

import "fmt"

type BjmInt int

type mapListInt = map[string][]int

type Person struct {
	name  string
	age   int
	hobby []string
	cp    *Person
}

func newPerson(name string) *Person {
	// 一系列操作...
	return &Person{
		name:  name,
		age:   20,
		hobby: make([]string, 5),
		cp:    nil,
	}
}

func test1() {
	pBenjamin := newPerson("Benjamin")
	pStella := newPerson("Stella")
	pBenjamin.cp = pStella
	pStella.cp = pBenjamin

	fmt.Printf("bjm: %p\n", pBenjamin)
	fmt.Printf("stella: %p\n", pStella)
	fmt.Printf("bjm.cp: %p\n", pBenjamin.cp)
	fmt.Printf("stella.cp: %p\n", pStella.cp)

	fmt.Println(pBenjamin.cp.name)
	fmt.Println(pStella.cp.name)
}

func test2() {
	s1 := ""
	for i := 0; i < 3; i++ {
		n, _ := fmt.Scan(&s1)
		fmt.Println(n)
	}

}

func main() {
	test2()
}
