//package main
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
//	queue := []*TreeNode{root}
//	ans := [][]int{}
//	for len(queue) > 0 {
//		tmp := []int{}
//		q := []*TreeNode{}
//		for _, node := range queue {
//			if node == nil {
//				continue
//			}
//			tmp = append(tmp, node.Val)
//
//			if node.Left != nil {
//				q = append(q, node.Left)
//			}
//			if node.Right != nil { // 记清楚root和node，别老是弄混
//				q = append(q, node.Right)
//			}
//		}
//		ans = append(ans, tmp)
//		queue = q
//	}
//	return ans
//}
