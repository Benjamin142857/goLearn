package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)
	fmt.Println(now.Date())
	fmt.Println(now.Clock())
	fmt.Println(now.Add(time.Hour * 24))

	timer := time.Tick(time.Second)
	for t := range timer {
		fmt.Println(t)
		if t.After(now.Add(time.Second * 10)) {
			break
		}
	}
}
