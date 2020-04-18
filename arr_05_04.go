package main

import "fmt"

func main() {

	a := [...]string{"AAA", "BBBB", "CCCC", "DDDD"}

	//这里是值传递,会拷贝一份数据
	b := a

	b[0] = "Frank"

	fmt.Println("a=", a)
	fmt.Println("b=", b)

	printArray("a", a)
	printArray("b", b)

}

// 格式化 数组 格式化信息
func printArray(name string, x [4]string) {

	fmt.Print(name, "\t")
	fmt.Printf("addr=%p \t len=%d \t cap=%d \t x=%v \n", &x, len(x), cap(x), x)

}
