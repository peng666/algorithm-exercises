package main

import "fmt"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		res = append(res, queue[len(queue)-1].Val)
		var nextLevel []*TreeNode
		for _, node := range queue {
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		queue = nextLevel
	}
	return res
}

func main() {
	node5 := &TreeNode{Val: 5}
	node4 := &TreeNode{Val: 4, Left: node5}
	node3 := &TreeNode{Val: 3}
	node2 := &TreeNode{Val: 2, Left: node4}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}
	res := rightSideView(node1)
	fmt.Print(res)
}
