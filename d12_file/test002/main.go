package main

import (
	"fmt"
	fileUtil "goLearn/d12_file/fileUtils"
)

func test1() {
	s, err := fileUtil.ReadAllByRead("../test001/test.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(s)
	}
}

func test2() {
	s, err := fileUtil.ReadAllByBufIO("../test001/test.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(s)
	}
}

func test3() {
	s, err := fileUtil.ReadAllByIOUtil("../test001/test.txt")
	if err != nil {
		panic(err)
	} else {
		fmt.Println(s)
	}
}

func test4() {
	var s string
	s += "我是Benjamin\n"
	s += "我爱Stella\n"
	s += "嘻嘻嘻嘻嘻嘻嘻666\n"
	err := fileUtil.WriteByWrite("myTest2.txt", s)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("WriteByWrite successful!!")
	}
}

func test5() {
	var s string
	s += "我是Benjamin\n"
	s += "我爱Stella\n"
	s += "嘻嘻嘻嘻嘻嘻嘻\n"
	err := fileUtil.WriteByBufIO("myTest3.text", s)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("WriteByBufIO successful!!")
	}

}

func test6() {
	var s string
	s += "我是Benjamin\n"
	s += "我爱Stella\n"
	s += "嘻嘻嘻嘻嘻嘻嘻\n"
	err := fileUtil.WriteByIOUtil("myTest4.text", s)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("WriteByIOUtil successful!!")
	}

}

func main() {
	test5()
}
