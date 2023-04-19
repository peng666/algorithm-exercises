//package main
//
//import (
//	"container/list"
//	"fmt"
//)
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	var res []int
//	//inorder(root, &res) // 递归
//	// 非递归。栈
//	stack := list.New()
//	for root != nil || stack.Len() > 0 {
//		for root != nil {
//			stack.PushFront(root)
//			root = root.Left
//		}
//
//		node := stack.Remove(stack.Front()).(*TreeNode)
//		res = append(res, node.Val)
//		root = node.Right
//	}
//	return res
//}
//
//func inorder(node *TreeNode, res *[]int) {
//	if node == nil {
//		return
//	}
//	inorder(node.Left, res)
//	*res = append(*res, node.Val)
//	inorder(node.Right, res)
//}
//
//func main() {
//	node1 := &TreeNode{Val: 1}
//	node2 := &TreeNode{Val: 2}
//	node3 := &TreeNode{Val: 3}
//	node4 := &TreeNode{Val: 4}
//	node1.Left = node2
//	node2.Right = node3
//	node2.Left = node4
//
//	res := inorderTraversal(node1)
//	fmt.Println(res)
//}
