package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var once sync.Once

func f1(ch chan int) {
	defer func() {
		wg.Done()
		fmt.Println("f1 END")
	}()

	for i := 0; i < 100; i++ {
		ch <- i
		ch <- i
		<-ch
	}
	once.Do(func() { close(ch) })
}

func f2(ch chan int) {
	defer func() {
		wg.Done()
		fmt.Println("f2 END")
	}()

	for {
		i, ok := <-ch
		fmt.Println(i, ok)
		if !ok {
			i, ok := <-ch
			fmt.Println(i, ok)
			break
		}
	}
}

func main() {
	ch := make(chan int, 100)
	wg.Add(2)
	go f1(ch)
	go f2(ch)
	wg.Wait()
	once.Do(func() { close(ch) })
	fmt.Println("main END")
}
