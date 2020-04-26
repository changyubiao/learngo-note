package main

import "fmt"

func main() {

	// 错误写法
	//var m map[string]int
	//m["frank"] = 1
	//fmt.Println(m)

	// 正确写法
	var m2 = make(map[string]int)
	m2["height"] = 165

	fmt.Println(m2)
}
