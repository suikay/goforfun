package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

type calFunc func(int, int) int

type Operation struct {
	Str string
	f   calFunc
}

var ops = []Operation{
	{
		Str: "+",
		f: func(x, y int) int {
			return x + y
		},
	},
	{
		Str: "-",
		f: func(x, y int) int {
			return x - y
		},
	},
	{
		Str: "×",
		f: func(x, y int) int {
			return x * y
		},
	},
	{
		Str: "÷",
		f: func(x, y int) int {
			return x / y
		},
	},
}

var congrats = []string{
	"好棒！",
	"恭喜你，正确了！",
	"你简直是天才的宝宝！",
}

var sads = []string{
	"差一点哦~",
	"继续努力~",
	"你再想想？",
}

func getRandSlice[T any](s []T) T {
	return s[rand.Intn(len(s))]
}

func main() {
	var cnt int
	rand.Seed(time.Now().Unix())

	color.HiCyan("勇敢的挑战者，请问你要挑战多少题？")
	fmt.Scanf("%d", &cnt)
	startTime := time.Now()
	right, wrong := 0, 0
	for i := 0; i < cnt; i++ {
		var ans int

		x, y, op := rand.Intn(100), rand.Intn(100), getRandSlice(ops)
		fmt.Printf("%d %s %d = ?\n", x, op.Str, y)
		fmt.Scanf("%d", &ans)
		if ans == op.f(x, y) {
			// fmt.Printf("%s\033[0m\n\n", getRandSlice(congrats))
			color.Green(getRandSlice(congrats) + "\n")
			right++
		} else {
			color.Red(getRandSlice(sads) + "\n")
			fmt.Printf("正确答案是：%d %s %d = %d\n\n",
				x, op.Str, y, op.f(x, y))
			wrong++
		}
	}
	fmt.Printf("************** 战绩 **************\n"+
		"用时：%.2f秒，每题平均用时：%.2f秒\n"+
		"正确：%d 错误： %d\n",
		time.Since(startTime).Seconds(), time.Since(startTime).Seconds()/float64(cnt), right, wrong)
}
