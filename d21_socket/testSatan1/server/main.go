package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var exitMutex sync.Mutex
var wg sync.WaitGroup
var exitFlag bool = false
var tempBufLen int = 1024

func preHandleRequest(conn net.Conn, sReq *string) {
	var buf []byte
	pReader := bufio.NewReader(conn)

	for {
		tempBuf := make([]byte, tempBufLen)
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
	*sReq = string(buf)
	return
}

func preHandleResponse(conn net.Conn, sRes string) {
	sResHeader := ""
	_, _ = conn.Write([]byte(sResHeader + sRes))
	return
}

func doProcess(conn net.Conn) {
	defer wg.Done()
	defer func() { _ = conn.Close() }()

	var sReq string
	var sRes string
	preHandleRequest(conn, &sReq)

	// do something

	fmt.Printf("[Worker] request input: %v\n", sReq)
	if sReq == "exit" {
		exitMutex.Lock()
		exitFlag = true
		exitMutex.Unlock()
	} else {
		sRes = strings.ToUpper(sReq)
	}

	preHandleResponse(conn, sRes)
}

func main() {
	var (
		url  = "127.0.0.1"
		port = "8899"
	)
	listen, err := net.Listen("tcp", url+":"+port)
	if err != nil {
		fmt.Println("net."+
			"Listen failed: ", err)
		return
	}
	defer func() { _ = listen.Close() }()

	fmt.Printf("[Main] Server Start on %v:%v >>>>>>>>>>>>>>>>>>>>>>>\n", url, port)
	for !exitFlag {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept failed: ", err)
			return
		}
		wg.Add(1)
		go doProcess(conn)
	}

	wg.Wait()
	fmt.Println("[Main] Server End >>>>>>>>>>>>>>>>>>>>>>>")
}
