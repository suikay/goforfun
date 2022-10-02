package main

import (
	"fmt"
	color "github.com/fatih/color"
	"math/rand"
	"time"
)

var opStr = []string{"+", "-", "*", "/"}

func opToStr(op int) string {
	return opStr[op]
}

func calAnswer(x, op, y int) int {
	switch op {
	case 0:
		return x + y
	case 1:
		return x - y
	case 2:
		return x * y
	case 3:
		return x / y
	}
	return 0
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

func getRandSlice(s []string) string {
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

		x, y, op := rand.Intn(100), rand.Intn(100), rand.Intn(2)
		fmt.Printf("%d %s %d = ?\n", x, opToStr(op), y)
		fmt.Scanf("%d", &ans)
		if ans == calAnswer(x, op, y) {
			// fmt.Printf("%s\033[0m\n\n", getRandSlice(congrats))
			color.Green(getRandSlice(congrats) + "\n")
			right++
		} else {
			color.Red(getRandSlice(sads) + "\n")
			fmt.Printf("正确答案是：%d %s %d = %d\n\n",
				x, opToStr(op), y, calAnswer(x, op, y))
			wrong++
		}
	}
	fmt.Printf("************** 战绩 **************\n"+
		"用时：%.2f秒，每题平均用时：%.2f秒\n"+
		"正确：%d 错误： %d\n",
		time.Since(startTime).Seconds(), time.Since(startTime).Seconds()/float64(cnt), right, wrong)
}
