package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg sync.WaitGroup
var mData sync.Map
var o sync.Once

func main() {
	for i := 0; i < 20; i++ {
		go func(x int) {
			wg.Add(1)
			kv := strconv.Itoa(x)
			mData.Store(kv, kv)
			v, ok := mData.Load(kv)
			if !ok {
				fmt.Println("error")
			}
			fmt.Printf("mData[%v]=%v\n", x, v)
			wg.Done()
		}(i)
	}

	wg.Wait()

	mData.Range(func(k, v interface{}) bool {
		fmt.Printf("k=%v, v=%v\n", k, v)
		return true
	})
}
