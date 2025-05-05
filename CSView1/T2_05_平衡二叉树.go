package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return abs(height(root.Left), height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0 // 写递归：记得写好终止条件！
	}
	return max(height(root.Left), height(root.Right)) + 1

}

func abs(x, y int) int {
	if x >= y {
		return x - y
	}
	return y - x
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
