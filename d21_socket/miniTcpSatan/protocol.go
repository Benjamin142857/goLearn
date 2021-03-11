package miniTcpSatan

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

type TcpPackage struct {
	msg string
}

// TcpPackage 构造函数
func NewTcpPackage(msg string) *TcpPackage {
	return &TcpPackage{
		msg: msg,
	}
}

// TcpPackage 编码
func (t *TcpPackage) Encode() (buf []byte, err error) {

	return
}

// TcpPackage 解码
func (t *TcpPackage) Decode(buf []byte) (err error) {
	net.Conn
	return
}
