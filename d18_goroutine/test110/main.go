package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var gNum []int = make([]int, 100000)
var wg sync.WaitGroup

func func1(x int) {
	defer wg.Done()
	gNum[x] = 1

	time.Sleep(time.Second)

	runtime.GOMAXPROCS()
}

func main() {
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func1(i)
	}
	wg.Wait()
	sum := 0
	for _, v := range gNum {
		sum += v
	}
	fmt.Println(sum)
	fmt.Println("main end")
}
