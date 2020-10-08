package main

import (
	"fmt"
)

// 数组是值类型 （赋值即深拷贝）
func test1() {
	a := [...]int{1, 2, 3, 4}
	b := a
	fmt.Println(a == b)
	b[3] = 100
	fmt.Println(a == b)
	fmt.Println(a)
	fmt.Println(b)
}

// 数组的遍历
func test2() {
	a := [...]string{"Benjamin", "love", "stella"}
	for i, v := range a {
		fmt.Println(i, v)
	}
}

// 二维数组
func test3() {
	a := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	b := [...][2]int{
		[2]int{1, 2},
	}
	fmt.Println(a)
	fmt.Println(b)
}

// 二维数组的遍历
func test4() {
	a := [...][3]string{
		[3]string{"Benjamin", "Stella", "Alex"},
		[3]string{"Love", "You", "Baby"},
		[3]string{"Benjamin", "Stella", "Alex"},
		[3]string{"Love", "You", "Baby"},
		[3]string{"Benjamin", "Stella", "Alex"},
		[3]string{"Love", "You", "Baby"},
		[3]string{"Benjamin", "Stella", "Alex"},
		[3]string{"Love", "You", "Baby"},
	}
	fmt.Printf("\n")
	for ix, x := range a {
		for iy, y := range x {
			fmt.Printf("[%v, %v] %s\t\t\t", ix, iy, y)
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
}

// 练习题1 - 求数组和
func practice1() int {
	arr := [...]int{1, 3, 5, 7, 8}
	cnt := 0
	for _, v := range arr {
		cnt += v
	}
	return cnt
}

// 练习题2 - 找出数组中和为指定值得两个元素的下标
func practice2(arr []int, target int) [][2]int {
	// 1. 冒泡排序
	flag := false
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
		flag = false
	}

	// 2. 双指针
	lp := 0
	rp := len(arr) - 1
	num := 0
	ret := make([][2]int, len(arr), len(arr))
	for lp < rp {
		vSum := arr[lp] + arr[rp]
		if vSum > target {
			rp--
		} else if vSum < target {
			lp++
		} else {
			ret[num] = [2]int{lp, rp}
			num++
			lp++
		}
	}
	return ret[:num]
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	res := practice2(arr, 10)
	fmt.Println(res)
}
