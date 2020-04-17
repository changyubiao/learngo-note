package main

import "fmt"

func main_01() {
	// [数组长度] 类型 {}
	var nums = [5]int{1, 2, 3, 4, 5}

	fmt.Println(nums)

	var nums2 = [...]int{2, 3, 5, 6, 7, 8}

	nums2[1] = 100
	fmt.Println(nums2)

}

func main_02() {
	a := [10]float64{1.9, 3, 6.7, 45, 34.2, 67.8}
	b := [...]float32{3, 4, 6.13}

	fmt.Println(a)

	fmt.Printf(" len(a) = %d, len(b)= %d", len(a), len(b))
}

func main_03() {

	arr := [5]int{1, 2, 4, 5, 6}

	for i, value := range arr {
		fmt.Println(i, value)
	}
}

func main() {
	var a = [4][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}

	fmt.Println("a=", a)

	fmt.Println(len(a))
	fmt.Println(len(a[0]))

	//	遍历二维数组
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}

}
