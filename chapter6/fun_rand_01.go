package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*

rand.Intn(3)   [0,3)  左闭右开区间


*/

func main() {
	//fmt.Println(rand.Int())
	//rand.Int()
	//
	//fmt.Println(rand.Float64())
	//rand.Float64()
	//fmt.Println(rand.Float32())

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Print(rand.Intn(3)," ")
	}
	fmt.Println()

	//randAnswer()
	//rand.Int()
	rand.Seed(time.Now().UnixNano())

	// 获取 0.0 - 1.0 的随机数
	fmt.Println(rand.Float64())


}

func getNumber() {
	rand.Seed(time.Now().UnixNano())
	// 获取 [0,10) 的随机数
	fmt.Println(rand.Intn(10))
}

func getNumberMN(m, n int) int {
	// 获取m-n 随机数
	rand.Seed(time.Now().UnixNano())

	num := rand.Intn(n-m+1) + m
	return num

}



func randTest() {
	//1、通过默认的随机数种子获取随机数.
	//系统默认的rand对象，随机种子默认都是1
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(50))
	fmt.Println(rand.Float64())

	//	2、动态随机种子，生成随机资源，实例化成随机对象，通过随机对象获取随机数
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	randnum := r1.Intn(10)
	fmt.Println(randnum)

	//3、简写形式：动态种子来获取随机数
	//	[0,10]
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(10))

	fmt.Println(rand.Float64())

	//[5,11]
	num := rand.Intn(7) + 5
	fmt.Println(num)
}



func randAnswer() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}

	// 设置 一个随机
	rand.Seed(time.Now().UnixNano())
	randnum := rand.Intn(len(answers))
	fmt.Println("随机回答", answers[randnum])
}
