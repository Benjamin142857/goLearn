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
	ctx1 := context.Background()
	ctx2, _ := context.WithCancel(ctx1)
	ctx3 := context.WithValue(ctx2, "main", "benjamin")

	fmt.Printf("ctx1: %v, %p\n", ctx1, ctx1)
	fmt.Printf("ctx2: %v, %p\n", ctx2, ctx2)
	fmt.Printf("ctx3: %v, %p\n", ctx3, ctx3)

	fmt.Println(ctx1.Value("main"))
	fmt.Println(ctx2.Value("main"))
	fmt.Println(ctx3.Value("main"))
}
