package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var stack []*TreeNode
	cur := root
	for len(stack) > 0 || cur != nil {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		if node.Right != nil {
			cur = node.Right // 这里要注意。不用加到队列，但指针要指对。指到右节点（还没进过队列的）
		}
	}
	return res
}

func inorderTraversal_digui(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	if root.Left != nil {
		left := inorderTraversal_digui(root.Left)
		res = append(res, left...)
	}
	res = append(res, root.Val)
	if root.Right != nil {
		right := inorderTraversal_digui(root.Right)
		res = append(res, right...)
	}
	return res
}

func main() {
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node3}
	node1 := &TreeNode{Val: 1, Right: node2}
	res := inorderTraversal(node1)
	fmt.Print(res)
}
