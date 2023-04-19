//package main
//
//import "fmt"
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	depth := 0
//	// 递归。
//	//depth = max(maxDepth(root.Left), maxDepth(root.Right)) + 1
//
//	// 非递归。队列。层次遍历。
//	queue := []*TreeNode{root}
//	for len(queue) > 0 {
//		depth++
//		q := []*TreeNode{}
//		for _, node := range queue {
//			if node.Left != nil {
//				q = append(q, node.Left)
//			}
//			if node.Right != nil {
//				q = append(q, node.Right)
//			}
//		}
//		queue = q
//	}
//	return depth
//}
//
//func max(num1, num2 int) int {
//	if num1 > num2 {
//		return num1
//	}
//	return num2
//}
//
//func main() {
//	node1 := &TreeNode{Val: 1}
//	node2 := &TreeNode{Val: 1}
//	node3 := &TreeNode{Val: 1}
//	node4 := &TreeNode{Val: 1}
//	node5 := &TreeNode{Val: 1}
//	node6 := &TreeNode{Val: 1}
//	node1.Left = node2
//	node2.Right = node3
//	node2.Left = node4
//	node4.Left = node5
//	node5.Right = node6
//
//	res := maxDepth(node1)
//	fmt.Println(res)
//}
