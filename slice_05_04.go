package main

import (
	"fmt"
	"strconv"
)

func main_08() {
	/*
		切片 是引用类型， 如果任何一个切片改变大小，会影响到其他其他的值
	*/

	a := [4]float32{1.0, 2.0, 3.0, 4.0}

	b := []int{2, 3, 4}

	fmt.Printf("&a =%p, 类型=%T, 数值 = %v, 长度= %d \n", &a, a, a, len(a))

	fmt.Printf("&b =%p, 类型=%T, 数值 = %v, 长度= %d \n", &b, b, b, len(b))

	c := a
	d := b
	fmt.Printf("&c =%p, 类型=%T, 数值 = %v, 长度= %d \n", &c, c, c, len(c))
	fmt.Printf("&d =%p, 类型=%T, 数值 = %v, 长度= %d \n", &d, d, d, len(d))

	a[1] = 200.001
	fmt.Println("a= ", a, ",c= ", c)

	c[3] = 300.001
	fmt.Println("a= ", a, ",c= ", c)

	b[0] = 200
	fmt.Println("b=", b, ",d=", d)

	d[1] = 300
	fmt.Println("b= ", b, "d= ", d)

}

func main_09() {
	/*

		修改 切片 数值， 多个切片共享 相同的数组的时候， 每个切片对数组的修改，都会体现在 数组数据的改变


	*/

	arr := [4]int{1, 2, 3, 4}

	nums1 := arr[:]
	nums2 := arr[:]

	fmt.Println("arr=", arr)

	printSlice("nums1", nums1)
	printSlice("nums2", nums2)

	nums1[0] = 100
	fmt.Println("arr=", arr) // 100, 2,3,4

	nums2[1] = 200
	fmt.Println("arr=", arr) // 100,200,3,4

}

// 格式化  切片 格式化信息
func printSlice(name string, x []int) {

	fmt.Print(name, "\t")
	fmt.Printf("addr:%p \t len=%d \t cap=%d \t slice=%v\n", x, len(x), cap(x), x)

}

func main_10() {

	numbers := make([]int, 0, 20)
	printSlice("numbes", numbers)

	numbers = append(numbers, 1)
	printSlice("numbes", numbers)

	numbers = append(numbers, 2)
	printSlice("numbes", numbers)

	numbers = append(numbers, 3, 4, 5)
	printSlice("numbes", numbers)

	s1 := []int{10, 11, 12}
	fmt.Printf("append s1: %v \n", s1)
	numbers = append(numbers, s1...)
	printSlice("numbers", numbers)

	fmt.Println("delete numbers[0] element ")
	numbers = numbers[1:]
	printSlice("numbers", numbers)

	a := int(len(numbers) / 2)

	fmt.Printf("a = %d,中间数 = %d\n", a, numbers[a])

	numbers = append(numbers[:a], numbers[a+1:]...)
	printSlice("numbers", numbers)

	fmt.Println("=== 创建一个切片numbers1  cap  是之前 切片的两倍")

	numbers1 := make([]int, len(numbers), 2*cap(numbers))

	printSlice("numbers1", numbers1)

	fmt.Println("copy from number1 to  numbers2")

	// copy 过来的 两个只是 数据，两个slice 没有任何关系，可以通过观察地址是不同的，只是拷贝了数据，注意
	count := copy(numbers1, numbers)
	fmt.Println("count=", count)

	printSlice("numbers1", numbers1)

	// 分别修改 两个slice 的一些值
	numbers[len(numbers)-1] = 666
	numbers1[0] = 111

	printSlice("numbers", numbers)
	printSlice("numbers1", numbers1)

}

func main() {

	var sa []string

	//sa := make([]string, 0, 20)

	printSliceMsg(sa)

	for i := 0; i < 15; i++ {
		sa = append(sa, strconv.Itoa(i))

		printSliceMsg(sa)
	}

}

func printSliceMsg(sa []string) {

	fmt.Printf("addr:%p \t len:%d \t cap:%d \t value:%v\n", sa, len(sa), cap(sa), sa)

}
