package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Print(pRoot *TreeNode) [][]int {
	if pRoot == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{pRoot}
	flag := false

	for len(queue) > 0 {
		tmp := []int{}
		q := []*TreeNode{}
		for _, node := range queue {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		if flag {
			reverse(&tmp)
		}
		res = append(res, tmp)
		queue = q
		flag = !flag
	}
	return res
}
func reverse(array *[]int) {
	n := len(*array) - 1
	for i := n / 2; i >= 0; i-- {
		tmp := (*array)[i]
		(*array)[i] = (*array)[n-i]
		(*array)[n-i] = tmp
	}
}

func main() {
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node1.Left = node2
	node1.Right = node3
	node3.Left = node5
	node2.Left = node4

	res := Print(node1)
	fmt.Println(res)
}
