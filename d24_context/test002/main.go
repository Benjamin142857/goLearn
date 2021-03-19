package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func f1(ctx context.Context) {
	defer wg.Done()
	wg.Add(1)
	go f2(ctx)
	fmt.Printf("f1: %p\n", ctx.Done())
}

func f2(ctx context.Context) {
	defer wg.Done()
	fmt.Printf("2: %p\n", ctx.Done())

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, "main", "benjamin")
	wg.Add(1)
	go f1(ctx)

	fmt.Printf("main: %p\n", ctx.Done())
	cancel()
	wg.Wait()
}
