package main

import (
	"fmt"
	"strings"
)

//func descMap(iter map[string]int, title string) {
//	fmt.Printf("[%p] %s: \t len=%v, cap=%v,   %v\n", iter, title, len(iter), cap(iter), iter)
//}

func descSlice(iter []int, title string) {
	fmt.Printf("[%p] %s: \t len=%v, cap=%v,   %v\n", iter, title, len(iter), cap(iter), iter)
}

func test1() {
	s1 := make([]map[string]string, 5, 10)
	fmt.Println(s1)
	s1[0] = make(map[string]string, 5)
	s1[0]["name"] = "Benjamin"
	fmt.Println(s1[0] == nil)
	fmt.Println(s1)
}

func test2() {
	m1 := make(map[string][]int, 10)
	fmt.Println(m1)
	fmt.Println(m1["Bjm"] == nil)
	descSlice(m1["Bjm"], "m1[Bjm]")
	//fmt.Println(m1["Bjm"])
	m1["Bjm"] = append(m1["Bjm"], 1, 2, 3)
	descSlice(m1["Bjm"], "m1[Bjm]")
	//m1["Benjamin"] =

}

// 练习：写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
func test3(str string) map[string]int {
	str = strings.Trim(str, " \n\t")
	str = strings.ToLower(str)
	s := strings.FieldsFunc(str, func(c rune) bool { return c == ' ' || c == ',' || c == '.' || c == '!' || c == '?' })
	m := make(map[string]int, len(s))
	for _, word := range s {
		m[word]++
	}
	return m
}

func test4() {
	type SliceInt []int
	s := make(SliceInt, 10)
	fmt.Println(s)
}

func test5() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s)
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s)
	fmt.Printf("%+v\n", m["q1mi"])
}

func main() {
	//testStr := `
	//	Hooray! It's snowing! It's time to make a snowman.James runs out. He makes a big pile of snow. He puts a big snowball on top. He adds a scarf and a hat. He adds an orange for the nose. He adds coal for the eyes and buttons.In the evening, James opens the door. What does he see? The snowman is moving! James invites him in. The snowman has never been inside a house. He says hello to the cat. He plays with paper towels.A moment later, the snowman takes James's hand and goes out.They go up, up, up into the air! They are flying! What a wonderful night!The next morning, James jumps out of bed. He runs to the door.He wants to thank the snowman. But he's gone.
	//`
	//m := test3(testStr)
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	test5()
}
