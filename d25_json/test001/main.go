package main

import (
	"encoding/json"
	"fmt"
)

type Color int8

const (
	red Color = iota
	white
	black
	gold
)

type SatanString string

type Dog struct {
	Name  string `json:"name"`
	Color Color  `json:"color"`
}

type Person struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Hobby []string `json:"hobby"`
	Dog   Dog      `json:"dog"`
}

func encode(i interface{}) string {
	s, err := json.Marshal(i)
	if err != nil {
		fmt.Println("json.Marshal fail")
		return ""
	}
	return string(s)
}

func main() {
	p1 := &Person{
		Name: "Benjamin",
		Age:  22,

		Hobby: []string{"coding", "loving"},
		Dog:   Dog{Name: "wangwang", Color: gold},
	}
	p2 := &Person{
		Name:  "Stella",
		Age:   22,
		Hobby: []string{"coding", "loving"},
		Dog:   Dog{Name: "wangwang", Color: gold},
	}

	res := make([]*Person, 0)
	fmt.Println(encode([]*Person{p1, p2}))

	err := json.Unmarshal([]byte(encode([]*Person{p1, p2})), &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(*res[0])
}
