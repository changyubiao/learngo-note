package main

import "fmt"

/*
数组指针 ： 首先是一个指针， 一个数组的地址

*[4]int
*[5]string

指针数组 ： 首先是一个数组，储存的数据类型是 指针
[4]*Type

[5]*int 数组 ,存储了5 int的地址的数组
[5]*float64  数组 存储了 5 个 float 地址的数组
[6]*string   数组   存储了6个字符串的指针地址的数组
[3]*string  数组  存储了3个字符串地址的数组
[3]string  数组  存储了3个字符串的数组


*[5]float64 , 指针，一个存储了5 个浮点类型数据点的数组的指针
*[3]string  , 指针，一个存储了3个 字符串的指针地址的数组

*[5]*float64 指针，一个数组指针， 存储了5个float64 类型的数据的指针地址的数组指针

*[3]*string  指针，存储了3个字符串的指针地址的数组的指针
**[4]string  指针，存储了4个字符串数据的数组的指针的指针

**[4]*string  指针，存储了4个字符串的指针地址的数组 的指针的指针

*/
func main() {
	//Pointer01()
	pointerArray()

}

func Pointer01() {
	// 1 创建一个数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("arr1=", arr1)

	// 2 创建一个指针，存储该数组的地址--->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)
	fmt.Printf("p1=%p, &p1=%p \n", p1, &p1)

	// 3 根据 数组指针 ，操作数组
	// 修改 数组的元素
	(*p1)[0] = 100
	fmt.Println("arr1=", arr1)

	// 简化写法
	p1[0] = 200
	fmt.Println("arr1=", arr1)

}

func pointerArray() {
	a := 1
	b := 2
	c := 3
	d := 4

	arr2 := [4]int{a, b, c, d}
	arr3 := [4]*int{&a, &b, &c, &d}
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)

	// 修改 arr2[0]
	arr2[0] = 100
	fmt.Println("arr2=", arr2)
	fmt.Println("a=", a)

	// 修改 arr3 的值
	//arr3[0] = 200
	*(arr3[0]) = 200
	fmt.Println("arr3=", arr3, ",a=", a)
	//fmt.Println("a=", a)

	a = 1
	//
	b = 1000
	fmt.Println("arr2=", arr2)
	fmt.Println("arr3=", arr3)

	for i := 0; i < len(arr3); i++ {
		fmt.Print(" ", *(arr3[i]))
	}
	fmt.Println()
}
