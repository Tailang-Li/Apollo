package leetcode

/*
Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
*/

func removeDuplicates(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	p := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[p] = nums[i]
			p++
		}
	}

	return p
}
