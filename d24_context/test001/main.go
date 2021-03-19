package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type current struct {
	stop bool
	info string
}

func f1(c *current) {
	defer func() { fmt.Println(c.info, " END") }()
	for !c.stop {
		fmt.Println(c.info)
		time.Sleep(800 * time.Millisecond)
	}
}

func main() {
	pReader := bufio.NewReader(os.Stdin)
	cList := make([]*current, 5)

	for i := 0; i < 5; i++ {
		c := &current{stop: false, info: fmt.Sprintf("goroutine_%v", i)}
		cList[i] = c
		go f1(c)
	}

	for {
		inputText, err := pReader.ReadString('\n')
		if err != nil {
			fmt.Println("无效输入, 请重新按要求输入...")
		} else {
			inputText = strings.TrimSpace(inputText)
			fmt.Printf("指令输入成功: %v\n", inputText)

			if inputText == "exit" {
				break
			} else if inputText == "0" {
				cList[0].stop = true
			} else if inputText == "1" {
				cList[1].stop = true
			} else if inputText == "2" {
				cList[2].stop = true
			} else if inputText == "3" {
				cList[3].stop = true
			} else if inputText == "4" {
				cList[4].stop = true
			}
		}
	}

}
