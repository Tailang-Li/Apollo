package leetcode

/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {

	for i := n + m - 1; i >= 0; i-- {
		if m == 0 || (n != 0 && nums2[n-1] > nums1[m-1]) {
			nums1[i] = nums2[n-1]
			n--
		} else {
			nums1[i] = nums1[m-1]
			m--
		}
	}
}
