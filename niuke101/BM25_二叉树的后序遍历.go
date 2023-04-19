//package main
//
//import "fmt"
//import "container/list"
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func postorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//
//	var res []int
//	//postorder(root, &res) // 递归
//	// 非递归。栈。
//	stack := list.New()
//	pre := &TreeNode{} // 用来记住访问过的父节点
//	for root != nil || stack.Len() > 0 {
//		for root != nil {
//			stack.PushFront(root)
//			root = root.Left
//		}
//		root = stack.Front().Value.(*TreeNode)
//		if root.Right != nil && root.Right != pre {
//			root = root.Right
//		} else {
//			res = append(res, root.Val)
//			stack.Remove(stack.Front())
//			pre = root
//			root = nil
//		}
//	}
//	return res
//}
//
//func postorder(node *TreeNode, res *[]int) {
//	if node == nil {
//		return
//	}
//	postorder(node.Left, res)
//	postorder(node.Right, res)
//	*res = append(*res, node.Val)
//}
//
//func main() {
//	node1 := &TreeNode{Val: 1}
//	node2 := &TreeNode{Val: 2}
//	node3 := &TreeNode{Val: 3}
//	node4 := &TreeNode{Val: 4}
//	node5 := &TreeNode{Val: 5}
//	node1.Left = node2
//	node2.Left = node3
//	node2.Right = node4
//	node4.Right = node5
//
//	res := postorderTraversal(node1)
//	fmt.Println(res)
//}
