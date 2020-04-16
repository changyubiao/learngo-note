package main

import "fmt"

func main_01() {
	/* 匿名函数 */
	func(data int) {
		fmt.Println("hello ", data)
	}(1000)
}

func main() {
	f := func(data string) {
		fmt.Println(data)
	}

	f("Welcome to study Golang!")
}
