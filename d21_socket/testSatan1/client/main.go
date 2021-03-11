package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

var exitMutex sync.Mutex
var exitFlag bool = false
var url = "127.0.0.1:8899"
var tempBufLen int = 1024

func test1() {
	//data := make(map[string]string, 10)
	//
	//conn, err := net.Dial("tcp", "127.0.0.1:8899")
	//if (err != nil) {
	//	fmt.Println("net.Dial failed")
	//	return
	//}
	//defer func () { _ = conn.Close() }()
	//
	//pStdinReader := bufio.NewReader(os.Stdin)

	//pStdinReader.ReadString("\n")
}

func preHandleRequest(conn net.Conn, sReq string) {
	_, _ = conn.Write([]byte(sReq))
	return
}

func preHandleResponse(conn net.Conn, sRes *string) {
	var buf []byte
	pReader := bufio.NewReader(conn)

	for {
		tempBuf := make([]byte, 128)
		n, err := pReader.Read(tempBuf)
		if err != nil {
			fmt.Println("pReader.Read failed.")
			return
		}

		buf = append(buf, tempBuf[:n]...)
		if n < tempBufLen {
			break
		}
	}
	*sRes = string(buf)
	return
}

func doRequest(data map[string]string, pReader *bufio.Reader) {
	fmt.Printf("Request: ")
	sData, err := pReader.ReadString('\n')
	sData = strings.TrimSpace(sData)
	if err != nil {
		fmt.Println("pReader.ReadString failed.")
		return
	}
	if sData == "quit" {
		exitMutex.Lock()
		exitFlag = true
		exitMutex.Unlock()
		return
	}

	data["data"] = sData
	conn, err := net.Dial("tcp", url)
	if err != nil {
		fmt.Println("net.Dial failed")
		return
	}
	defer func() { _ = conn.Close() }()

	// map => str
	var sReq, sRes string
	for k, v := range data {
		sReq += fmt.Sprintf("%v=%v&", k, v)
	}
	sReq = sReq[:len(sReq)-1]

	preHandleRequest(conn, sReq)
	preHandleResponse(conn, &sRes)
	fmt.Printf("Response: %v\n", sRes)
}

func test2() {
	// 0. 预声明
	fmt.Println("Client Start >>>>>>>>>>>>>>>>>>>")
	pStdinReader := bufio.NewReader(os.Stdin)
	data := make(map[string]string, 10)

	//	1. 先输入用户名
	fmt.Printf("input username: ")
	username, err := pStdinReader.ReadString('\n')
	username = strings.TrimSpace(username)
	if err != nil {
		fmt.Println("pReader.ReadString failed.")
		return
	}
	data["username"] = username

	// 2. 循环请求
	for !exitFlag {
		doRequest(data, pStdinReader)
	}

	// 3. 结束
	fmt.Println("Client END >>>>>>>>>>>>>>>>>>>")
}

func main() {
	test2()
}
