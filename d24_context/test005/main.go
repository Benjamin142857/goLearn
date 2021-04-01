// 测试一个 ctx.Done() channel 到底可以拿多少个消息
package main

import (
	"context"
	"fmt"
	"time"
)

func doing(name string, ctx context.Context) {
	for {
		x, ok := <-ctx.Done()
		fmt.Printf("child[%v] x=[%v]...\n", name, x)
		fmt.Printf("child[%v] ok=[%v]...\n", name, ok)
		if !ok {
			break
		}
	}

	//for x, ok := range ctx.Done() {
	//	fmt.Printf("child[%v] x=[%v]...\n", name, x)
	//	fmt.Printf("child[%v] ok=[%v]...\n", name, ok)
	//
	//	fmt.Printf("%v recv Done...\n", name)
	//}
}

func child1(ctx context.Context) {
	fmt.Println("child1 start >>>>>>>>>>>>")
	doing("child1", ctx)
	fmt.Println("child1 end >>>>>>>>>>>>")
}

func main() {
	ctx1, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	go child1(ctx1)
	time.Sleep(time.Second * 10)

	fmt.Println("main end")
}
