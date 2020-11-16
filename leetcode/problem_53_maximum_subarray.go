package leetcode

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Follow up: If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

func maxSubArray(nums []int) int {

	max, sa := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		sa = maxNum(sa+nums[i], nums[i])
		if sa > max {
			max = sa
		}
	}

	return max
}

func maxNum(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
