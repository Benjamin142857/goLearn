package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func A() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("A: %v\n", i)
	}
}

func B() {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		fmt.Printf("B: %v\n", i)
	}
}

func test1() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go A()
	go B()
	wg.Wait()
}

func main() {
	test1()
}
