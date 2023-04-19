//package main
//
//import (
//	"fmt"
//)
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//
//	// 非递归。队列。
//	res := [][]int{}
//	// slice版队列
//	queue := []*TreeNode{root}
//	for len(queue) > 0 {
//		tmp := []int{}
//		q := []*TreeNode{}
//		for _, node := range queue {
//			tmp = append(tmp, node.Val)
//			if node.Left != nil {
//				q = append(q, node.Left)
//			}
//			if node.Right != nil {
//				q = append(q, node.Right)
//			}
//		}
//		res = append(res, tmp)
//		queue = q
//	}
//	// list版的队列
//	//m := 1
//	//queue := list.New()
//	//queue.PushFront(root)
//	//for queue.Len() > 0 {
//	//	tmp := []int{}
//	//	n := m
//	//	m = 0
//	//	for ; n > 0; n-- {
//	//		node := queue.Remove(queue.Back()).(*TreeNode)
//	//		tmp = append(tmp, node.Val)
//	//		if node.Left != nil {
//	//			queue.PushFront(node.Left)
//	//			m++
//	//		}
//	//		if node.Right != nil {
//	//			queue.PushFront(node.Right)
//	//			m++
//	//		}
//	//	}
//	//	res = append(res, tmp)
//	//}
//	return res
//}
//
//func main() {
//	node1 := &TreeNode{Val: 1}
//	node2 := &TreeNode{Val: 2}
//	node3 := &TreeNode{Val: 3}
//	node4 := &TreeNode{Val: 4}
//	node5 := &TreeNode{Val: 5}
//	node6 := &TreeNode{Val: 6}
//	node1.Left = node2
//	node1.Right = node3
//	node2.Right = node4
//	node4.Right = node5
//	node4.Left = node6
//
//	res := levelOrder(node1)
//	fmt.Print(res)
//}
