package main

import "fmt"

type Animal struct {
	feet int
	name string
}

func (a *Animal) walk() {
	fmt.Printf("Walk by %v feet.\n", a.feet)
}

type Dog struct {
	Animal
	size int
}

type Bird struct {
	Animal
	winkColor string
}

type theWalk interface {
	walk()
}

func catch(w theWalk) {
	fmt.Printf("%T\n", w)
	w.walk()
}

func printAllType(elms ...interface{}) {
	for _, v := range elms {
		fmt.Printf("%T\n", v)
	}
}

func main() {
	a1 := &Animal{}
	a1.name = "Alex"
	a1.feet = 8

	d1 := &Dog{}
	d1.name = "WangWang"
	d1.feet = 4
	d1.size = 100

	b1 := &Bird{}
	b1.name = "Stella"
	b1.feet = 2
	b1.winkColor = "Pink"

	catch(a1)
	catch(b1)
	catch(d1)

	printAllType(a1, b1, d1, []int{1, 2, 3}, map[string]int{"aa": 123}, 30, "kkk")
}
