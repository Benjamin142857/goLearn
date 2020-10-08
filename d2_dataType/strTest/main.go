package main

import (
	"fmt"
	"strings"
)

func main() {
	nameBoy := "Benjamin嘿嘿"
	nameGirl := "Stella哈哈"
	they := nameBoy + nameGirl
	fmt.Println(they)
	fmt.Println("===============\n\n")

	fieldString := "aaa=0|int|require\nbbb=1|long|require\nccc=2|string|optional"
	fieldString = strings.Trim(fieldString, " \n")
	fieldList := strings.Split(fieldString, "\n")
	for i, v := range fieldList {
		fmt.Println(i, v)
		fmt.Println(strings.Contains(v, "require"))
	}
	//fmt.Println(fieldList)
	fmt.Println("===============\n\n")

	//for i, v := range they {
	//	fmt.Println(i, byte(v))
	//}
	testStr := []byte(nameGirl)
	//fmt.Println(testStr)
	testStr[8] = 128
	fmt.Println(testStr)
	fmt.Println(string(testStr))

}
