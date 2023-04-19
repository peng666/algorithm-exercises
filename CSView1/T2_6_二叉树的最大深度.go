//package CSView1
//
//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//// 迭代。层次遍历
//func maxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	queue := []*TreeNode{root}
//	ans := 0
//	for len(queue) > 0 {
//		count := len(queue)
//		for count > 0 {
//			node := queue[0]
//			queue = queue[1:]
//			if node.Left != nil {
//				queue = append(queue, node.Left)
//			}
//			if node.Right != nil {
//				queue = append(queue, node.Right)
//			}
//			count--
//		}
//		ans++
//	}
//	return ans
//}
