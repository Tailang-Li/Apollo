package leetcode

/*
Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	idx := len(nums) / 2

	return &TreeNode{
		Val:   nums[idx],
		Left:  sortedArrayToBST(nums[:idx]),
		Right: sortedArrayToBST(nums[idx+1:]),
	}
}
