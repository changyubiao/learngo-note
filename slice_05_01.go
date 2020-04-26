package main

import "fmt"

/*

类型 ,长度, 容量
make([]T , length , capacity)



array vs  slice  的区别?

数组是一种具有相同类型的一组长度固定的数据序列，数组的内存是一段连续的存储区域，定义后大小不能更改。
Slice（切片）代表边可变长度的序列，序列中每个元素都是相同的类型。Slice和数组的语法很像。
从底层来看，切片引用了数组的对象。切片可以追加元素，在追加时可能使切片的容量增大。
切片与数组相比，不需要设定长度，在[]中不用设定值，相对来说比较自由。


*/

func main_01() {
	// 定义切片
	var nums []int
	var nums2 []string = make([]string, 5)

	num3 := make([]string, 5)

	fmt.Println(nums, nums2, num3)
}

func main_02() {
	// 定义切片
	var nums = make([]int, 3, 5)
	fmt.Printf("%T \n", nums)

	fmt.Printf("len=%d, cap=%d,slice=%v \n", len(nums), cap(nums), nums)

	//  初始化 切片

	s := []int{1, 2, 4}

	arr := [10]int{1, 2, 3, 5, 6, 7, 8, 9}
	s2 := arr[:]

	s3 := arr[3:6]

	fmt.Printf("len=%d, cap=%d,s=%v \n", len(s), cap(s), s)
	fmt.Printf("len=%d, cap=%d,s2=%v \n", len(s2), cap(s2), s2)
	fmt.Printf("len=%d, cap=%d,s3=%v \n", len(s3), cap(s3), s3)

}

func sliceCap() {

	/*
		slice 本身没有存储任何数据，共享底层数组的数据
	*/
	arr := [...]string{"a", "b", "c", "d", "ee", "ff", "gg", "hh", "ii", "jj", "kk"}
	fmt.Printf("len=%d, cap=%d,s3=%v \n", len(arr), cap(arr), arr)

	// 左闭右开区间
	s0 := arr[2:9]
	fmt.Printf("%T,cap=%d , s0=%v \n", s0, cap(s0), s0)

	s2 := arr[6:8]
	fmt.Printf("%T, cap=%d ,s2=%v \n", s2, cap(s2), s2)

	s3 := s0[1:3]
	fmt.Printf("%T, cap=%d ,s3=%v \n", s2, cap(s3), s3)

	// 切片是引用类型，直接影响到 其他的 slice 的情况
	s3[0] = "frank"

	s4 := arr[1:5]
	fmt.Printf("%T, cap=%d ,s4=%v \n", s4, cap(s4), s4)

	fmt.Println(arr)
	fmt.Println(s0, s2, s3, s4)
}

func main() {
	sliceCap()
}
