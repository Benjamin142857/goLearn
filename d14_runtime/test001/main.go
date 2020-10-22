package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
)

type CallInfo struct {
	fileName string
	lineNo   int
	funcName string
}

func getInfo(skip int) (info CallInfo, err error) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		err = os.ErrInvalid
		return
	}
	pFunc := runtime.FuncForPC(pc)

	info.fileName = path.Base(file)
	info.lineNo = line
	info.funcName = pFunc.Name()
	return
}

func test1() {
	var info CallInfo

	info, _ = getInfo(0)
	fmt.Println(info)

	info, _ = getInfo(1)
	fmt.Println(info)

	info, _ = getInfo(2)
	fmt.Println(info)
}

func main() {
	test1()
}
