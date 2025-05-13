package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	for root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		curLevel := make([]int, 0, len(queue))
		var nextQueue []*TreeNode
		for i := 0; i < len(queue); i++ {
			node := queue[i]
			curLevel = append(curLevel, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		queue = nextQueue
		res = append(res, curLevel)
	}

	return res
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}

	node1.Left = node2
	node1.Right = node3
	node3.Right = node4
	node4.Left = node5

	res := levelOrder(node1)
	fmt.Println(res)
}
