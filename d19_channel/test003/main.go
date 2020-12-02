package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(id int, chJob <-chan int, chRes chan<- int) {
	defer wg.Done()

	chRes <- 1
	for jobs := range chJob {
		time.Sleep(time.Millisecond * 100)
		fmt.Printf("[%v], get Jobs %v.\n", id, jobs)

		if jobs == 72057594037927935 {
			close(chRes)
		} else {
			chRes <- jobs*2 + 1
		}
	}
}

func main() {
	//wg.Add(2)
	//ch1 := make(chan int, 1)
	//ch2 := make(chan int, 1)
	//go worker(1, ch1, ch2)
	//go worker(2, ch2, ch1)
	//wg.Wait()

	fmt.Println(123 / 10)
}
