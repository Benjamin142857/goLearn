package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var doCounter = [4]int{0, 0, 0}

func worker(id int, jobsChan <-chan int) {
	defer wg.Done()

	for job := range jobsChan {
		time.Sleep(time.Millisecond * 30)
		fmt.Printf("[%v] handle a job: %v\n", id, job)
	}
}

func main() {
	wg.Add(3)

	// run worker
	ch1 := make(chan int, 5)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 1)
	go worker(1, ch1)
	go worker(2, ch2)
	go worker(3, ch3)

	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 10)
		select {
		case ch1 <- i:
			doCounter[0]++
		case ch2 <- i:
			doCounter[1]++
		case ch3 <- i:
			doCounter[2]++
		default:
			doCounter[3]++
			fmt.Println("All worker busy, pass...")
		}
	}
	close(ch1)
	close(ch2)
	close(ch3)

	wg.Wait()
	for idx, doCnt := range doCounter {
		fmt.Printf("[%v] do jobs = %v\n", idx, doCnt)
	}
}
