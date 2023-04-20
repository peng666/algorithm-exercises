//package CSView1
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
//	// 非递归。栈。
//	stack := []*TreeNode{root}
//	res := []int{}
//	for len(stack) > 0 {
//		node := stack[len(stack)-1] // 出栈
//		stack = stack[:len(stack)-1]
//		if node != nil {
//			res = append(res, node.Val)
//		}
//		if node.Right != nil { // 想清楚哈！是右节点先入栈！！！
//			stack = append(stack, node.Right)
//		}
//		if node.Left != nil {
//			stack = append(stack, node.Left)
//		}
//	}
//	return res
//}
