package main

import "fmt"

func test1() {
	name := ""
	age := 0
	fmt.Printf("请输入您的姓名与年龄: ")
	n, err := fmt.Scan(&name, &age)
	fmt.Println("您的姓名是: ", name, ", 年龄是: ", age)
	fmt.Println(n, err)
	fmt.Print()
}

func test2() {
	name := ""
	age := 0
	fmt.Println("请按照规范输入字符串: ")
	fmt.Scanf("name:%s age:%d", &name, &age)
	fmt.Println(name, age)
}

func test3() {
	var (
		name    = "fangzhiwei"
		age     = 22
		company = "huya"
	)
	str := fmt.Sprintf("name: %v, age: %v, company: %v", name, age, company)
	fmt.Println(str)
}

func main() {
	test3()
}
