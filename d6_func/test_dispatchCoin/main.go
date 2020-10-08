package main

import (
	"fmt"
	"strings"
)

func dispatchCoin(coins int, users []string) (int, map[string]int) {
	rule := map[rune]int{
		'e': 1,
		'i': 2,
		'o': 3,
		'u': 4,
	}
	left := coins
	distribution := make(map[string]int, len(users))
	for _, name := range users {
		cnt := 0
		for _, c := range strings.ToLower(name) {
			coinNum, ok := rule[c]
			if ok {
				cnt += coinNum
			}
		}
		left -= cnt
		distribution[name] = cnt
	}

	return left, distribution
}

func main() {
	coins := 500
	users := []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter",
		"Giana", "Adriano", "Aaron", "Elizabeth",
	}
	left, distribution := dispatchCoin(coins, users)
	fmt.Println("left: ", left, "\n")
	for name, coin := range distribution {
		fmt.Println(name, coin)
	}
}
