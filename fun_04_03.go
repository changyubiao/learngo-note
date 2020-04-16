package main

import "fmt"

type processFun func(int) bool // 声明一个函数类型

func main() {

	fmt.Println("hello  func 40 03")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("nums=", nums)

	even := filter(nums, isEven)
	fmt.Println(even)
	fmt.Println(filter(nums, isOdd))
}

func isEven(num int) bool {
	/*
		判断是否为偶数*/
	if num%2 == 0 {
		return true
	} else {
		return false
	}
}

func isOdd(num int) bool {
	/* 判断是否为奇数
	 */

	if num%2 == 1 {
		return true
	} else {
		return false
	}

}

func filter(slice []int, f processFun) []int {

	var result []int

	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result

}
