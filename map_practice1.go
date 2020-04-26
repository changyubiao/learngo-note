package main

/*
(1)实现杨辉三角输出（打印10行）。
*/

import "fmt"

func main() {
	var arr [10][10]int //定义一个能容纳10行的杨辉三角

	for i := 0; i < 10; i++ {
		arr[i][0] = 1 //定义第一列都为1
		arr[i][i] = 1 //定义对角线都为1
	}
	for i := 2; i < 10; i++ {
		for j := 1; j < 10; j++ {
			arr[i][j] = arr[i-1][j-1] + arr[i-1][j] //利用上一行递推计算
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10-i; j++ {
			fmt.Print(" ") //打印每一行的空格，使之成为三角形
		}
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}

}
