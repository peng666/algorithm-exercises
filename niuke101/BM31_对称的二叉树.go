package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetrical(root *TreeNode) bool {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return true
	}

	return help(root.Left, root.Right) // 递归。变形的前序遍历。
	// 非递归。2个队列。

}
func help(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return help(root1.Left, root2.Right) && help(root2.Left, root1.Right)
}
