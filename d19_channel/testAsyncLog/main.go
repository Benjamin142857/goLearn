package main

import (
	"fmt"
	"goLearn/d19_channel/bjmLogAsync"
	"os"
	"time"
)

func test1() {
	logLevel := bjmLogAsync.DEBUG
	pFile, _ := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	logger := bjmLogAsync.NewLogger(pFile, &logLevel)

	go logger.Run()

	var err error

	err = logger.Info("哈哈哈哈这是个info")
	if err != nil {
		fmt.Println(err)
	}
	err = logger.Info("哈哈哈哈这是个info")
	if err != nil {
		fmt.Println(err)
	}
	err = logger.Info("哈哈哈哈这是个info")
	if err != nil {
		fmt.Println(err)
	}
	logger.Debug("哈哈哈哈这是个debug")
	if err != nil {
		fmt.Println(err)
	}
	logger.Debug("哈哈哈哈\n这是\n个warn")
	if err != nil {
		fmt.Println(err)
	}
	logger.Error("哈哈哈哈\n这是\n个error\n")
	if err != nil {
		fmt.Println(err)
	}

	logger.Stop()

	logger.Debug("哈哈哈哈\n这是\n个warn")
	if err != nil {
		fmt.Println(err)
	}
	logger.Error("哈哈哈哈\n这是\n个error\n")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("log END")

	time.Sleep(time.Second * 10)
}

func main() {
	test1()
	fmt.Println("Main END")
}
