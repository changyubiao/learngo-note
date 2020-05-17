package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main01() {

	fmt.Println("请输入 用户名 密码：")
	username :="frank"
	pwd :="password"
	fmt.Scanln(&username,&pwd)

	fmt.Println("账号信息",username,pwd)



}



func main() {
	playGame()
}

//游戏过程
func playGame() {
	//获取随机数
	target := generateRandNum(10, 100)
	fmt.Println("请输入随机数[10,100]：")
	fmt.Println("--------------------")

	//记录猜测的次数
	count := 0
	for {
		count++
		yourNum := 0
		fmt.Scanln(&yourNum)

		if yourNum < target {
			fmt.Println("小了❌")
		} else if yourNum > target {
			fmt.Println("大了❌")
		} else {
			fmt.Println("正确✅")
			fmt.Printf("您一共猜测了%d次！ \n", count)
			fmt.Println("--------------------")
			playGame()
		}
		//错误提示
		alertInfo(count , target)
	}

}

//错误提示
func alertInfo(count, target int) {
	if count >=6  {
		fmt.Printf("您一共猜了 %d 次都没有猜中，太笨了！😓 " , count)
		fmt.Println("正确数字：" , target)
		fmt.Println("--------------------")
		playGame()
	}
}

//生成随机数
func generateRandNum(min, max int) int {
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
