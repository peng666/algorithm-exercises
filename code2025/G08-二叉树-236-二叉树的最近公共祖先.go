package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil // 这个子树找不到
	}
	if root == p || root == q {
		return root // 这个子树找到了
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root // 两边都找着了，祖先就是当前节点
	}
	if left != nil {
		return left // right没找着，说明都在left那边，left返回来的就是祖先
	}
	return right
}

func main() {
	node7 := &TreeNode{Val: 7}
	node4 := &TreeNode{Val: 4}
	node6 := &TreeNode{Val: 6}
	node2 := &TreeNode{Val: 2, Left: node7, Right: node4}
	node0 := &TreeNode{Val: 0}
	node8 := &TreeNode{Val: 8}
	node5 := &TreeNode{Val: 5, Left: node6, Right: node2}
	node1 := &TreeNode{Val: 1, Left: node0, Right: node8}
	node3 := &TreeNode{Val: 3, Left: node5, Right: node1}

	res := lowestCommonAncestor(node3, node5, node4)
	fmt.Println(res.Val)
}
