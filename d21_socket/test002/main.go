package main

import (
	"fmt"
	"math"
	"net"
)

func test1() {
	ch := make(chan int)

	x := <-ch
	fmt.Println(x)
}

func encode(strLen int32) (lenBuff [4]byte) {
	for idx, _ := range []int{1, 2, 3, 4} {
		lenBuff[3-idx] = byte(int8(strLen % 255))
		strLen /= 255
	}
	return
}

func decode(lenBuff [4]byte) (strLen int32) {
	strLen = 0
	for idx, bt := range lenBuff {
		strLen += int32(float64(bt) * math.Pow(255, float64(3-idx)))
	}
	return
}

func test001() {
	lb := encode(1234)
	fmt.Println(lb)
	sl := decode(lb)
	fmt.Println(sl)
}

func test002() {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		fmt.Println(iface.Flags)
		fmt.Println(net.FlagUp)
		fmt.Println(net.FlagLoopback)
		fmt.Println(iface.Addrs())
	}
}

func main() {
	test002()
}
