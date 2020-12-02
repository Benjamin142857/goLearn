package main

import (
	"fmt"
	"math"
	"time"
)

type TaskItem struct {
	idx   int
	value int
}

func newTaskItem(idx int, value int) *TaskItem {
	return &TaskItem{
		idx:   idx,
		value: value,
	}
}

func producer(ch chan<- *TaskItem) {
	for i := 0; i < 100; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Printf("producer|idx=%v|value=%v\n", i, i)
		ch <- newTaskItem(i, i)
	}
	close(ch)
	fmt.Println("producer END")
}

func consumer(ch <-chan *TaskItem, chRes chan<- *TaskItem) {
	for pTask := range ch {
		time.Sleep(time.Millisecond * 500)
		pTask.value = int(math.Pow(float64(pTask.value), 2)) + 1
		chRes <- pTask
	}
	close(chRes)
	fmt.Println("consumer END")
}

func main() {
	ch := make(chan *TaskItem, 100)
	chRes := make(chan *TaskItem, 1)

	go producer(ch)
	go consumer(ch, chRes)

	for pTaskRes := range chRes {
		fmt.Printf("consumer|idx=%v|value=%v\n", pTaskRes.idx, pTaskRes.value)
	}
	fmt.Println("main END")
}
