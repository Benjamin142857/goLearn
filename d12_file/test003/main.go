package main

import (
	"fmt"
	"goLearn/d12_file/bjmLog"
	"os"
	"path"
)

func test1() {
	lL := bjmLog.INFO
	logger := bjmLog.NewLogger(os.Stdout, &lL)
	logger.Info("哈哈哈哈这是个info")
	logger.Info("哈哈哈哈这是个info")
	logger.Info("哈哈哈哈这是个info")
	logger.Debug("哈哈哈哈这是个debug")
	logger.Warn("哈哈哈哈\n这是\n个warn")
	logger.Error("哈哈哈哈\n这是\n个error\n")
	fmt.Println("--------------------------------")
	lL = bjmLog.ERROR
	logger.Info("哈哈哈哈这是个info")
	logger.Info("哈哈哈哈这是个info")
	logger.Info("哈哈哈哈这是个info")
	logger.Debug("哈哈哈哈这是个debug")
	logger.Warn("哈哈哈哈\n这是\n个warn")
	logger.Error("哈哈哈哈\n这是\n个error\n")
}

func test2(lL *bjmLog.LogLevel) {
	pFile, _ := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	logger := bjmLog.NewLogger(pFile, lL)
	logger.Info("哈哈哈哈这是个info")
	logger.Info("哈哈哈哈这是个info")
	logger.Info("哈哈哈哈这是个info")
	logger.Debug("哈哈哈哈这是个debug")
	logger.Debug("哈哈哈哈\n这是\n个warn")
	logger.Error("哈哈哈哈\n这是\n个error\n")
}

func test3() {
	fmt.Println(path.Dir("A/B///C/D"))
	fmt.Println(path.Base("A/B//C/D/"))

}

func test4() {
	fmt.Printf("%T, %v\n", bjmLog.DEBUG, bjmLog.DEBUG)
	fmt.Printf("%T, %v\n", bjmLog.WARN, bjmLog.WARN)
}

func main() {
	test1()
}
