package main

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context) {
	fmt.Println("child start >>>>>>>>>>>>")
	for {
		select {
		case x := <-ctx.Done():
			fmt.Printf("child: type(x)=%t, x=%v\n", x, x)
			goto childEnd
		default:
			fmt.Println("child: doing...")
		}
		time.Sleep(time.Millisecond * 500)
	}

childEnd:
	fmt.Println("child end >>>>>>>>>>>>")
}

func main() {
	fmt.Println("main start >>>>>>>>>>>>")
	ctx, cancelFunc := context.WithCancel(context.Background())
	go child(ctx)

	time.Sleep(time.Second * 5)
	fmt.Println("main cancel >>>>>>>>>>>>")
	cancelFunc()

	time.Sleep(time.Second * 5)
	fmt.Println("main end >>>>>>>>>>>>")

}
