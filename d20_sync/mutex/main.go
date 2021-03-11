package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	var v = 0

	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			defer wg.Done()
			mutex.Lock()
			v = v + 1
			mutex.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println(v)
}
