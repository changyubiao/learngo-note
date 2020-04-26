package main

/*
(2)写一个函数，判断一个字符串是否对称，若对称，返回true，否则返回false。

*/

import (
	"fmt"
)

func main() {
	fmt.Println("请输入字符串")
	var str string
	fmt.Scanln(&str)
	flag := isD(str)
	if flag == true {
		fmt.Println(str + "是对称字符串")
	} else {
		fmt.Println(str, "不是对称的字符串")
	}
}
func isD(str string) bool {
	arr := []byte(str)
	j := len(arr) - 1
	i := 0
	for i = 0; i < len(arr)-1; i++ {
		if arr[i] == arr[j] {
			j--
		} else {
			return false
		}
	}
	return true
}
