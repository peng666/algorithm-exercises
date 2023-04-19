//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func Convert(root *TreeNode) *TreeNode {
//	if root == nil || (root.Left == nil && root.Right == nil) {
//		return root
//	}
//	res := &TreeNode{}
//	pre := res
//	// 中序遍历。非递归。栈。
//	stack := []*TreeNode{}
//
//	for len(stack) > 0 || root != nil {
//		for root != nil {
//			stack = append(stack, root)
//			root = root.Left
//		}
//		node := stack[len(stack)-1]
//		stack = stack[0 : len(stack)-1]
//		val := node.Val
//		tmp := &TreeNode{Val: val, Left: pre, Right: nil}
//		pre.Right = tmp
//		pre = pre.Right
//		root = node.Right
//	}
//	res.Right.Left = nil
//	return res.Right
//}
//
//func main() {
//	node1 := &TreeNode{Val: 1}
//	node3 := &TreeNode{Val: 3}
//	node2 := &TreeNode{Val: 2}
//	node1.Left = node2
//	node1.Right = node3
//	Convert(node1)
//
//}
