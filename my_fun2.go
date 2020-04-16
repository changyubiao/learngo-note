package main

import "fmt"

/*
全局变量 和局部变量


*/

var num int = 10
var num2 int = 20

func main() {

	/* main 函数中 声明局部变量 */
	num, num2, height := 1, 2, 3

	fmt.Printf("man() num=%d , num2=%d,height=%d \n", num, num2, height)

	ret := sum(num, num2)
	fmt.Printf("main() ret=%d", ret)

}

func sum(num, num2 int) (ret int) {

	num++
	num2 += 2
	ret = num + num2

	fmt.Printf("sum()  num=%d,num2 =%d , ret =%d \n", num, num2, ret)
	return ret

}
