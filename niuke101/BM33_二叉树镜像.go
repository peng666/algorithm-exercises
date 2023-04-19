package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Mirror(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	left := Mirror(root.Left)
	right := Mirror(root.Right)
	root.Left = right
	root.Right = left
	return root
}
