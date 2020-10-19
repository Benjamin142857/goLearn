package main

import (
	"fmt"
	"os"
)

func test0() {
	pFile, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file error!")
		return
	}
	defer func() {
		err := pFile.Close()
		if err != nil {
			fmt.Println("close file error!")
		}
	}()

	fmt.Println(pFile)
}

func test1() {
	var (
		buff    []byte
		content string
	)

	pFile, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("open file error!")
		return
	}
	defer func() {
		err := pFile.Close()
		if err != nil {
			fmt.Println("close file error!")
		}
	}()

	buff = make([]byte, 128)
	for {
		n, err := pFile.Read(buff)
		if err != nil {
			fmt.Println("read file error!")
			return
		}
		fmt.Println(n)
		content += string(buff[:n-1])
		if n < 128 {
			break
		}
	}

	fmt.Println(content)
}

func main() {
	test0()
}
