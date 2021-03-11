package main

import (
	"net"
)

func main() {
	aa, _ := net.Listen("tcp", "127.0.0.1:8899")
	aa.Accept()
}
