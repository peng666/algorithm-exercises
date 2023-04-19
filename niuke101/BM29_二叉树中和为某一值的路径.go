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
//func hasPathSum(root *TreeNode, sum int) bool {
//	if root == nil {
//		return false
//	}
//
//	// 递归
//	//if root.Val == sum && root.Left == nil && root.Right == nil {
//	//	return true
//	//}
//	//return hasPathSum(root.Left, sum-root.Val) || hasPathSum(root.Right, sum-root.Val)
//
//	// 非递归。两个栈。深度优先搜索遍历dfs
//	stack1 := list.New() // 一个栈记录节点
//	stack2 := list.New() // 一个栈记录前缀和
//	stack1.PushFront(root)
//	stack2.PushFront(root.Val)
//	for stack1.Len() > 0 {
//		node := stack1.Remove(stack1.Front()).(*TreeNode)
//		val := stack2.Remove(stack2.Front()).(int)
//		if sum == val && node.Left == nil && node.Right == nil {
//			return true
//		}
//		if node.Left != nil {
//			stack1.PushFront(node.Left)
//			stack2.PushFront(node.Left.Val + val)
//		}
//		if node.Right != nil {
//			stack1.PushFront(node.Right)
//			stack2.PushFront(node.Right.Val + val)
//		}
//	}
//	return false
//}
//
//func main() {
//	node1 := &TreeNode{Val: 5}
//	node2 := &TreeNode{Val: 4}
//	node3 := &TreeNode{Val: 8}
//	node4 := &TreeNode{Val: 1}
//	node5 := &TreeNode{Val: 11}
//	node6 := &TreeNode{Val: 9}
//	node7 := &TreeNode{Val: 2}
//	node8 := &TreeNode{Val: 7}
//	node1.Left = node2
//	node1.Right = node3
//	node2.Left = node4
//	node2.Right = node5
//	node3.Right = node6
//	node5.Left = node7
//	node5.Right = node8
//
//	res := hasPathSum(node1, 22)
//	fmt.Println(res)
//}
