package main

import "fmt"

type aer interface {
	a()
}

type ber interface {
	a()
	b()
}

type cer interface {
	aer
	ber
}

type Person struct {
	name string
	age  int
}

func (p Person) callName() {
	fmt.Printf("name: %v\n", p.name)
}

func (p *Person) callAge() {
	fmt.Printf("age: %v\n", p.age)
	p.age++
}

func main() {

	p1 := Person{
		"Benjamin",
		21,
	}
	p2 := &Person{
		"Stella",
		22,
	}

	p1.callName()
	p2.callName()

	p1.callAge()
	p1.callAge()
	p2.callAge()
	p2.callAge()

}
