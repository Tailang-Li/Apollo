package main

import (
	"fmt"
)

/*
Given an integer array nums, return the length of the longest strictly increasing subsequence.

A subsequence is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, [3,6,2,7] is a subsequence of the array [0,3,1,6,2,2,7].
*/

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
	fmt.Println(lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7}))
	fmt.Println(lengthOfLIS([]int{0, 1, 0, 3, 2, 3}))
}

func lengthOfLIS(nums []int) int {

	var totalMax int
	cnts := make([]int, len(nums))
	for i := range nums {

		maxCnt := 1
		for j := i - 1; j >= 0; j-- {
			if nums[i] > nums[j] {
				maxCnt = max(maxCnt, cnts[j]+1)
				if cnts[j] == totalMax {
					break
				}
			}
		}

		totalMax = max(totalMax, maxCnt)
		cnts[i] = maxCnt
	}

	return totalMax
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
