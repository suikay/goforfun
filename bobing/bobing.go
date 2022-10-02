package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func isNothing(num []int) bool {
	if len(num) != 6 {
		return true
	}
	for _, x := range num {
		if x == 4 {
			return false
		}
	}
	if num[0] == num[3] || num[1] == num[4] || num[2] == num[5] {
		return false
	}
	return true
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var name string
	for j := 1; j <= 10; j = j + 1 {
		fmt.Printf("你的名字：")
		fmt.Scanln(&name)

		num := make([]int, 0)

		for i := 1; i <= 6; i = i + 1 {
			num = append(num, rand.Intn(6)+1)
		}
		sort.Slice(num, func(i, j int) bool { return num[i] < num[j] })

		fmt.Printf("%s: %v", name, num)

		if isNothing(num) {
			fmt.Printf("  乞丐，哈哈哈哈哈哈哈~~\n")
		} else {
			fmt.Printf("  手气不错~~\n")
		}
	}
}
