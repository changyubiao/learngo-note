package main

import "fmt"

/*
指针作为参数:

引用传递
值传递



其实本质上来说 都是值传递，
传递指针 其实也是传递一个值，只是这个值 是一个地址而已。





数组是值类型 ， 直接拷贝一份数据
*/

func main() {

	a := 1
	var arr = [4]int{1, 2, 3, 4}
	fun1(a)
	fmt.Println("fun1() 调用后,a=", a)

	fun2(&a)
	fmt.Println("fun2() 调用后,a=", a)

	fun3(arr)
	fmt.Println("fun3() 调用后,arr=", arr)

	fun4(&arr)
	fmt.Println("fun4() 调用后,arr=", arr)

	s1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println("s1=", s1)

}

func fun4(p *[4]int) {
	p[0] = 10000
	fmt.Println("fun3() 函数中修改arr[0]:", p[0])

}

func fun3(arr [4]int) {
	arr[0] = 10000
	fmt.Println("fun3() 函数中修改arr[0]:", arr[0])

}

func fun1(num int) {
	fmt.Println("fun1() 函数中num:", num)
	num = 100
	fmt.Println("fun1() 函数 中修改num:", num)
}

func fun2(p *int) {
	fmt.Println("fun2() 函数中num:", *p)
	*p = 100
	fmt.Println("fun2()函数中 修改num:", *p)

}
