package main

import "fmt"

type Person struct {
	School
	name string
	age  int
}

type School struct {
	name  string
	level int
}

func main() {
	p1 := Person{
		School{
			"GDUT",
			5,
		},
		"Benjamin",
		20,
	}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.School)
}
