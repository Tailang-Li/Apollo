package leetcode

/*
Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	levelMap map[int][]int
)

func levelOrderBottom(root *TreeNode) [][]int {

	levelMap = make(map[int][]int)
	levelOrder(root, 0)
	result := make([][]int, len(levelMap))
	for i := range result {
		result[i] = levelMap[len(result)-i-1]
	}

	return result
}

func levelOrder(root *TreeNode, level int) {

	if root == nil {
		return
	}

	levelOrder(root.Left, level+1)
	levelOrder(root.Right, level+1)

	if _, ok := levelMap[level]; !ok {
		levelMap[level] = make([]int, 0)
	}

	levelMap[level] = append(levelMap[level], root.Val)
}
