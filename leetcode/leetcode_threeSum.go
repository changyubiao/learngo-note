package main

import (
	"fmt"
	"sort"
)

func threeSum01(nums []int) [][]int {

	numsLen := len(nums)
	if numsLen < 3 {
		return nil
	}

	sort.Ints(nums)
	target := 0
	var lo, hi int
	var res [][]int

	for i := 0; i < numsLen-2; i++ {
		if nums[i] > target { // 如果最小的数大于0则都不满足条件
			continue
		}

		fmt.Println("i=", i)

		// 第一个数去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		lo = i + 1
		hi = numsLen - 1

		for lo < hi {
			// 后面的 位置去重
			if hi < numsLen-1 && nums[hi] == nums[hi+1] {
				hi--
				continue
			}

			left := nums[lo]
			right := nums[hi]
			sum := nums[i] + left + right

			//
			if sum < target {
				lo++
			} else if sum > target {
				hi--
			} else {
				// 相等的情况
				res = append(res, []int{nums[i], left, right})
				lo++
				hi--
			}
		}
	}
	return res
}

func main() {

	//var nums = []int{-1, 0, 1, 2, 3, -1, -4}
	//var nums = []int{-1, 0, 1, 2, -1, -4}
	var nums = []int{-2, 0, 1, 1, 2}

	r := threeSum01(nums)

	//fmt.Println(nums)
	fmt.Println(r)
}
