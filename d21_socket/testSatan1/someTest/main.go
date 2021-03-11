package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func preHandleRequest(r io.Reader, sReq *string) {
	var buf []byte
	pReader := bufio.NewReader(r)

	for {
		tempBuf := make([]byte, 128)
		n, err := pReader.Read(tempBuf)
		if err != nil {
			break
		}
		buf = append(buf, tempBuf[:n]...)
	}
	*sReq = string(buf)
	return
}

func test1() {
	pFile, err := os.Open("./aaa.txt")
	if err != nil {
		fmt.Println("os.Open failed.")
		return
	}
	defer func() { _ = pFile.Close() }()

	s := ""
	preHandleRequest(pFile, &s)
	fmt.Println(s)
}

func test2() {
	s1 := fmt.Sprintf("aaa: %v&", 1)
	fmt.Print(s1)
	s1 += fmt.Sprintf("bbb: %v&", 2)
	fmt.Print(s1)
	s1 += fmt.Sprintf("ccc: %v&", 3)
	fmt.Println("")
	fmt.Print(s1[:len(s1)-1])
}

func main() {
	test2()
}
