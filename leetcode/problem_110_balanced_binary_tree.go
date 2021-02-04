package leetcode

/*
Given a binary tree, determine if it is height-balanced.

For this problem, a height-balanced binary tree is defined as:

a binary tree in which the left and right subtrees of every node differ in height by no more than 1.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, isBalanced := height(root)
	return isBalanced
}

func height(root *TreeNode) (int, bool) {

	if root == nil {
		return 0, true
	}

	leftHeight, leftIsBalanced := height(root.Left)
	rightHeight, rightIsBalanced := height(root.Right)

	if !leftIsBalanced || !rightIsBalanced || abs(leftHeight, rightHeight) > 1 {
		return 0, false
	}

	return max(leftHeight, rightHeight) + 1, true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
