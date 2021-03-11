package main

import (
	"fmt"
	"strconv"
	"sync"
)

var mData = make(map[string]string, 1)
var wg sync.WaitGroup
var lock sync.RWMutex
var wLock sync.Mutex

func main() {
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(x int) {
			lock.Lock()
			mData["key"] = strconv.Itoa(x)
			lock.Unlock()
			lock.RLock()
			fmt.Printf("[%v] mData[key]=%v\n", x, mData["key"])
			lock.RUnlock()
			wg.Done()
		}(i)
	}

	wg.Wait()
}
