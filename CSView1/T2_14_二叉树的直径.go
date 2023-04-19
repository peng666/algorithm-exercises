//package CSView1
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func diameterOfBinaryTree(root *TreeNode) int {
//	res := 0
//	var helper func(*TreeNode) int
//	helper = func(root *TreeNode) int {
//		if root == nil {
//			return 0
//		}
//
//		left := helper(root.Left)
//		right := helper(root.Right)
//		res = max(res, left+right) // 想清楚！很简单的：有多少个子节点=就是有多少个路径
//		return max(left, right) + 1
//	}
//	helper(root)
//	return res
//}
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
