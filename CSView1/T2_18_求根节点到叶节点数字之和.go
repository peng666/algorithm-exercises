//package main
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func sumNumbers(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//
//	curSum := 0 // 外部变量用来做累加
//	var dfs func(*TreeNode, int)
//	dfs = func(node *TreeNode, curNum int) {
//		if node == nil {
//			return
//		}
//		curNum = curNum*10 + node.Val
//		if node.Left == nil && node.Right == nil {
//			curSum += curNum
//			return
//		}
//		dfs(node.Left, curNum)
//		dfs(node.Right, curNum)
//	}
//	dfs(root, 0) // 一开始传个0进去就可以了
//	return curSum
//}
