//package main
//
//import (
//	"container/list"
//	"fmt"
//)
//
///**
// * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
// *
// *
// * @param root TreeNode类
// * @return int整型一维数组
// */
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func preorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	var res []int
//	//preorder(root, &res)  // 递归
//
//	stack := list.New()
//	stack.PushFront(root)
//	//stack := []*TreeNode{root}
//	for stack.Len() > 0 {
//		node := stack.Remove(stack.Front()).(*TreeNode)
//		res = append(res, node.Val)
//		if node.Right != nil {
//			stack.PushFront(node.Right)
//		}
//		if node.Left != nil {
//			stack.PushFront(node.Left)
//		}
//	}
//	return res
//}
//
//func preorder(node *TreeNode, res *[]int) {
//	if node == nil {
//		return
//	}
//	*res = append(*res, node.Val)
//	preorder(node.Left, res)
//	preorder(node.Right, res)
//}
//func main() {
//	node1 := &TreeNode{Val: 1}
//	node2 := &TreeNode{Val: 2}
//	node3 := &TreeNode{Val: 3}
//	node4 := &TreeNode{Val: 4}
//	node1.Right = node2
//	node2.Left = node3
//	node3.Left = node4
//
//	res := preorderTraversal(node1)
//	fmt.Println(res)
//
//}
