package leetcode

/*
Given a sorted array of distinct integers and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
*/

func searchInsert(nums []int, target int) int {

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	var l, r int = 0, len(nums) - 1
	for l != r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return r
}
