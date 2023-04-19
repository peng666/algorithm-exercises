//package CSView1
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func isValidBST(root *TreeNode) bool {
//	return traverse(root, nil, nil)
//}
//
//func traverse(root, min, max *TreeNode) bool {
//	if root == nil {
//		return true
//	}
//	val := root.Val
//
//	if min != nil && val <= min.Val { // 记住这个递归条件~~~想清楚啦
//		return false
//	}
//	if max != nil && val >= max.Val {
//		return false
//	}
//	return traverse(root.Left, min, root) && traverse(root.Right, root, max)
//}
