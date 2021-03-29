package main

import (
	"fmt"
	"math"
)

/*
Given an array of n integers nums, a 132 pattern is a subsequence of three integers nums[i], nums[j] and nums[k] such that i < j < k and nums[i] < nums[k] < nums[j].

Return true if there is a 132 pattern in nums, otherwise, return false.

Follow up: The O(n^2) is trivial, could you come up with the O(n logn) or the O(n) solution?
*/

func main() {
	fmt.Println(find132pattern([]int{1, 2, 3, 4}))
	fmt.Println(find132pattern([]int{3, 1, 4, 2}))
	fmt.Println(find132pattern([]int{-1, 3, 2, 0}))
	fmt.Println(find132pattern([]int{-2, 1, 1}))
}

// wrong
func find132pattern(nums []int) bool {

	n := len(nums)
	if n < 3 {
		return false
	}

	secondaryMin := math.MinInt64
	stack := make([]int, 0)
	stack = append(stack, nums[0])

	for i := 2; i < len(nums); i++ {

		for len(stack) != 0 && nums[i-1] > stack[len(stack)-1] {
			secondaryMin = min(secondaryMin, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 || nums[i-1] < stack[len(stack)-1] {
			stack = append(stack, nums[i-1])
		}

		if nums[i] < stack[len(stack)-1] && nums[i] > secondaryMin {
			return true
		}
	}

	return false
}

func find132pattern2(nums []int) bool {

	n := len(nums)
	if n < 3 {
		return false
	}

	// 采用从左向右遍历 1 的方式，需要维护一个单调栈，用于维护一个最大的值，以及次大的值
	secondaryMax := ^(int(^uint(0) >> 1))

	// 初始化栈
	stack := make([]int, 0)
	stack = append(stack, nums[n-1])

	for i := n - 3; i >= 0; i-- {

		for len(stack) != 0 && nums[i+1] > stack[len(stack)-1] {
			secondaryMax = max(secondaryMax, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 || nums[i+1] != stack[len(stack)-1] {
			stack = append(stack, nums[i+1])
		}

		if nums[i] < secondaryMax {
			return true
		}
	}

	return false
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
