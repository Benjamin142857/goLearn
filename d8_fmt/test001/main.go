package main

import "fmt"

func test1() {
	fmt.Printf("%9.2f\n", 3313.545)
}

func test2() {
	fmt.Printf("%-.2s", "ABC")
	fmt.Print("666")
}

func main() {
	test2()
}
