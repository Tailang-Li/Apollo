package main

import "fmt"

/*
Given an array nums with n integers, your task is to check if it could become non-decreasing by modifying at most one element.

We define an array is non-decreasing if nums[i] <= nums[i + 1] holds for every i (0-based) such that (0 <= i <= n - 2).
*/

func main() {
	fmt.Println(checkPossibility([]int{3, 4, 2, 3}))
	fmt.Println(checkPossibility([]int{4, 2, 3}))
}

func checkPossibility(nums []int) bool {

	var count int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {

			count++
			if count > 1 {
				return false
			}

			if i != 0 && nums[i+1] < nums[i-1] {
				nums[i+1] = nums[i]
			}
		}
	}

	return true
}
