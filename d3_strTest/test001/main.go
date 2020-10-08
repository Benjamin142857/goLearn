package main

import "fmt"

func main() {
	strTest := "Hello沙河小王子Benjamin大帅比"
	var cnt int = 0
	for _, v := range strTest {
		if v > 128 {
			cnt++
		}
	}
	fmt.Println(cnt)

}
