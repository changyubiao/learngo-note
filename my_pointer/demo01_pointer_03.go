package main

import "fmt"

/*

函数 指针 ： 是一个指针， 指向一个函数的只做真
因为 golang 语言中 ，function ，默认看做一个指针 没有 *

引用类型

slice  map   function

指针 函数  ： 一个函数  这个函数返回一个指针



*/

func main() {
	// 定义一个函数类型
	var f func()
	f = fun1
	f()

	arr1 := fun2()
	fmt.Printf("arr1 类型: %T,地址: %p,数值: %v\n", arr1, &arr1, arr1)

	arr2 := fun3()
	fmt.Printf("arr2 类型: %T,地址: %p,数值: %v\n", arr2, &arr2, arr2)
	fmt.Printf("arr2的值(arr2中存储的数组地址): %p\n", arr2)

}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("fun3()中 arr类型: %T,地址: %p,数值: %v\n", arr, &arr, arr)
	return &arr
}

func fun2() [4]int {

	arr := [4]int{1, 2, 3, 4}

	return arr

}

func fun1() {

	fmt.Println("fun1() ....")

}
