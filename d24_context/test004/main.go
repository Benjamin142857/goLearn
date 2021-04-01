// 测试： 父 ctx.Done() 后 子 ctx也会Done吗

package main

import (
	"context"
	"fmt"
	"time"
)

func doing(name string, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("child[%v] doing...\n", name)
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func child1(ctx context.Context) {
	fmt.Println("child1 start >>>>>>>>>>>>")
	ctx2, _ := context.WithCancel(ctx)
	go child2(ctx2)
	doing("child1", ctx)
	fmt.Println("child1 end >>>>>>>>>>>>")
}

func child2(ctx context.Context) {
	fmt.Println("child2 start >>>>>>>>>>>>")
	doing("child2", ctx)
	fmt.Println("child2 end >>>>>>>>>>>>")
}

func main() {
	fmt.Println("main start >>>>>>>>>>>>")
	ctx, cancelFunc := context.WithCancel(context.Background())
	go child1(ctx)

	time.Sleep(time.Second * 5)
	fmt.Println("main cancel >>>>>>>>>>>>")
	cancelFunc()

	time.Sleep(time.Second * 5)
	fmt.Println("main end >>>>>>>>>>>>")
}
