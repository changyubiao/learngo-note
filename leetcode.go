package main

import "fmt"

//* Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}
	left := 0
	right := len(nums) - 1

	return build(nums, left, right)
}

func build01(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	var maxVal, index int

	//maxVal = -10000
	maxVal = nums[left]
	index = left
	fmt.Printf("maxvalue:%d ,left:%d\n", maxVal, left)

	for i := left; i <= right; i++ {
		if maxVal < nums[i] {
			index = i
			maxVal = nums[i]
		}

	}

	root := &TreeNode{Val: maxVal}
	root.Left = build(nums, left, index-1)
	root.Right = build(nums, index+1, right)

	return root
}

func build(nums []int, left int, right int) *TreeNode {

	if left > right {
		return nil
	}

	index := getMaxIndex(nums, left, right)
	root := &TreeNode{Val: nums[index]}
	root.Left = build(nums, left, index-1)
	root.Right = build(nums, index+1, right)

	return root
}

func getMaxIndex(nums []int, left int, right int) int {
	maxIndex := left
	for i := left; i <= right; i++ {
		if nums[maxIndex] < nums[i] {
			maxIndex = i
		}
	}
	return maxIndex
}

func main() {
	//var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	//make()
	//nums := [3,2,1,6,0,5]
	nums := [6]int{3, 2, 1, 6, 0, 5}
	fmt.Println(nums)

}
