package main

import "fmt"

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

	p := TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSymmetric(&p))
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return isSubSymmetric(root.Left, root.Right)
}

func isSubSymmetric(l *TreeNode, r *TreeNode) bool {

	if l == nil && r == nil {
		return true
	}

	if l == nil || r == nil || l.Val != r.Val {
		return false
	}

	return isSubSymmetric(l.Left, r.Right) && isSubSymmetric(l.Right, r.Left)
}
