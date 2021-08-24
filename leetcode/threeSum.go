package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	numsLen := len(nums)
	if numsLen < 3 {
		return nil
	}
	target := 0
	var lo, hi int
	var res [][]int
	for i := 0; i < numsLen-2; i++ {
		lo = i + 1
		hi = numsLen - 1
		if nums[i] > 0 { // 如果最小的数大于0则都不满足条件
			return nil
		}
		fmt.Println("i=", i)

		// 第一个数去重
		if i > 0 && nums[i] == nums[i+1] {
			continue
		}
		for lo < hi {
			fmt.Println("low", lo, " , hi ", hi)
			left := nums[lo]
			right := nums[hi]
			sum := nums[i] + left + right

			if sum < target {
				for lo < hi && nums[lo] == nums[lo+1] {
					lo++
				}
			} else if sum > target {
				for lo < hi && nums[hi] == nums[hi-1] {
					hi--
				}

			} else {
				res = append(res, []int{nums[i], left, right})

				//lo++
				//hi--
				for lo < hi && nums[lo] == left {
					lo++
				}
				for lo < hi && nums[hi] == right {
					hi--
				}
			}
		}
	}
	return res
}

func main() {

	var nums = []int{-1, 0, 1, 2, -1, -4}

	r := threeSum(nums)

	//fmt.Println(nums)
	fmt.Println(r)

}
