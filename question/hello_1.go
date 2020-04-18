package main

import (
	"fmt"
	"strconv"
)

/*
strconv.Itoa()函数的参数是一个整型数字，它可以将数字转换成对应的字符串类型的数字。


string函数的参数若是一个整型数字，它将该整型数字转换成ASCII码值等于该整形数字的字符。



*/
func main() {
	number := 97
	result := strconv.Itoa(number)

	//fmt.Println(result)
	fmt.Printf("%T ,result=%v\n", result, result)

	result2 := string(number)
	//fmt.Println(result2)
	fmt.Printf("%T ,result2=%v\n", result2, result2)

}
