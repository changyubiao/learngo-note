package main

import "fmt"

func main() {
	var a int = 10
	fmt.Printf("a:%v,变量a的地址:%p\n", a, &a)

	var p1 *int
	p1 = &a
	fmt.Println("p1的值:", p1)
	fmt.Printf("p1的地址:%p\n", &p1)
	fmt.Println("p1的数值，是a的地址，改地址存储的数据:", *p1)

	// 3.操作变量 更改值 ，并不会改变地址
	a = 100
	fmt.Println("a:", a)
	fmt.Printf("%p\n", &a)

	// 4 通过指针修改变量的值
	*p1 = 200
	fmt.Printf("4,  a:%v,变量a的地址:%p\n", a, &a)

	// 5 指针的指针
	var p2 **int
	fmt.Println("5, p2=", p2)

	p2 = &p1
	fmt.Printf("%T,%T,%T\n", a, p1, p2)

	fmt.Println("p2的数值：", p2)
	fmt.Printf("p2 自己的地址: %p\n", &p2)
	fmt.Println("p2 中存储的地址，对应的数值，就是p1的地址，对应的数据:", *p2)
	fmt.Println("p2中存储的地址，对应的数值，在获取对应的数值:", **p2)
}
