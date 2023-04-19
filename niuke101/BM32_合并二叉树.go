package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTree(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	head := &TreeNode{Val: root1.Val + root2.Val}
	head.Left = mergeTree(root1.Left, root2.Left)
	head.Right = mergeTree(root1.Right, root2.Right)
	return head
}
